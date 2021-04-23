#!/bin/bash

# Script to create a kind + sriov + ovnkube setup. Inspired by https://github.com/kubevirt/kubevirt/blob/master/cluster-up/cluster/kind-k8s-sriov-1.17.0/config_sriov.sh
set -exuo pipefail

SCRIPT=`realpath -s $0`
VDPA=${VDPA:-no}
SCRIPTPATH=`dirname $SCRIPT`
MANIFESTS_DIR=${SCRIPTPATH}/../deployment/ovnkube-sriov
. $SCRIPTPATH/common.sh


deploy_multus() {
  echo 'Deploying Multus'
  _kubectl create -f $MANIFESTS_DIR/multus.yaml
  _kubectl create -f $MANIFESTS_DIR/netattach.yaml

  echo 'Waiting for Multus deployment to become ready'
  daemonset_name=$(cat $MANIFESTS_DIR/multus.yaml | grep -i daemonset -A 3 | grep -Po '(?<=name:) \S*amd64$')
  daemonset_namespace=$(cat $MANIFESTS_DIR/multus.yaml | grep -i daemonset -A 3 | grep -Po '(?<=namespace:) \S*$' | head -1)
  required_replicas=$(_kubectl get daemonset $daemonset_name -n $daemonset_namespace -o jsonpath='{.status.desiredNumberScheduled}')
  wait_for_daemonSet $daemonset_name $daemonset_namespace $required_replicas

  return 0
}

deploy_deviceplugin() {
  echo 'Deploying DevicePlugin'
   if [ "${VDPA} " != "no" ]; then
  	_kubectl create -f $MANIFESTS_DIR/configMap-vdpa.yaml
   else
  	_kubectl create -f $MANIFESTS_DIR/configMap.yaml
   fi

  _kubectl create -f $MANIFESTS_DIR/sriov_device_plugin.yaml

  echo 'Waiting for DevicePlugin deployment to become ready'
  daemonset_name=$(cat $MANIFESTS_DIR/sriov_device_plugin.yaml | grep -i daemonset -A 3 | grep -Po '(?<=name:) \S*amd64$')
  daemonset_namespace=$(cat $MANIFESTS_DIR/sriov_device_plugin.yaml | grep -i daemonset -A 3 | grep -Po '(?<=namespace:) \S*$' | head -1)
  required_replicas=$(_kubectl get daemonset $daemonset_name -n $daemonset_namespace -o jsonpath='{.status.desiredNumberScheduled}')
  wait_for_daemonSet $daemonset_name $daemonset_namespace $required_replicas

  return 0
}

setup_sriov() {
        NUM_VFS=4
	local -r pf=$1

        echo 0 | sudo tee /sys/class/net/$pf/device/sriov_numvfs
        echo $NUM_VFS | sudo tee /sys/class/net/$pf/device/sriov_numvfs

        for i in $(seq 0 $(($NUM_VFS -1))); do
            echo "Unbinding VF ${i}"
            pci_addr=$(get_pci_addr ${pf} $i)
            echo $pci_addr | sudo tee /sys/bus/pci/drivers/mlx5_core/unbind || true
        done
	set_eswitch_mode $pf switchdev
	# Just in case, if NetworkManager is running, tell it not to handle the PF (or it will mess the qdiscs)
        sleep 5
	nmcli device set $pf managed no || true
	set_tc_offload ${pf}

        for i in $(seq 0 $(($NUM_VFS -1))); do
            echo "Binding VF ${i}"
            pci_addr=$(get_pci_addr ${pf} $i)
            echo $pci_addr | sudo tee /sys/bus/pci/drivers/mlx5_core/bind
            sleep 3
            devname=$(ls -x /sys/class/net/${pf}/device/virtfn${i}/net)
            [ -z "$devname" ] && error "Cannot get VF network device"
            set_tc_offload ${devname}
            nmcli device set ${devname} managed no || true
        done

        for i in $(seq 0 $(($NUM_VFS -1))); do
            devname=$(ls -x /sys/class/net/${pf}/device/virtfn${i}/net)
            rep=$(get_representor ${pf} $i)
            echo "Configuring ${devname} (rep: ${rep} )"
            set_tc_offload ${rep}
            nmcli device set ${rep} managed no || true
            sudo ip link set ${rep} up
        done
	sudo ip link set ${pf} up
}

# $1 is the vdpa device
# $2 is the desired vdpa driver
set_vdpa_driver() {
        local dev=$1
        local driver=$2
            if [ -d "/sys/bus/vdpa/devices/${dev}/driver" ]; then
                    local curr_driver=$(basename $(readlink /sys/bus/vdpa/devices/${dev}/driver))
                    if [[ "${curr_driver}" != "${driver}" ]]; then
                        echo "${dev}" | sudo tee /sys/bus/vdpa/drivers/${curr_driver}/unbind
                else 
                        return 0
                fi
        fi
        echo "${dev}" | sudo tee /sys/bus/vdpa/drivers/${driver}/bind
}

usage(){
    echo "Usage: $0 command"
    echo ""
    echo "commands:"
    echo "  start PFNAME:   Start the environment"
    echo "  stop:           Stop the environment"
    exit 1
}

do_start() {
    local -r pf=$1
    echo "Configuring SR-IOV"
    #docker run --rm --cap-add=SYS_RAWIO quay.io/phoracek/lspci@sha256:0f3cacf7098202ef284308c64e3fc0ba441871a846022bb87d65ff130c79adb1 sh -c "lspci | egrep -i 'network|ethernet'"

    echo "PFNAME: $pf"

    if [ "${VDPA} " != "no" ]; then
	    sudo modprobe vdpa || true
	    sudo modprobe virtio-vdpa || true
	    sudo modprobe vhost-vdpa || true
	    sudo modprobe mlx5-vdpa || true
    else
	    sudo modprobe -r mlx5-vdpa || true
	    sudo modprobe -r vhost-vdpa || true
	    sudo modprobe -r virtio-vdpa || true
	    sudo modprobe -r vdpa || true
    fi

    setup_sriov $pf

    if [ "${VDPA} " != "no" ]; then
	echo ""
	echo "Binding half of the devices to vhost-vdpa and other half to virtio-vpda"
	DRIVERS=("vhost_vdpa" "virtio_vdpa")
	devices=( $(ls -x /sys/bus/vdpa/devices) )
	for i in ${!devices[@]}; do
	    dev=${devices[$i]}
	    driver=${DRIVERS[$(($i%2))]}
	    echo "Binding device ${dev} to driver ${driver}"
	    set_vdpa_driver $dev $driver
	done
    fi

    deploy_multus 

    # TODO: sriov-operator
    # TODO: apply sriov_node_policy see: https://github.com/kubevirt/kubevirt/blob/master/cluster-up/cluster/kind-k8s-sriov-1.17.0/config_sriov.sh
    # Alternatively, deploy Device Plugin
    deploy_deviceplugin

    if [ -f "/etc/openvswitch/conf.db" ]; then
    	echo "About to remove old OpenvSwitch configuration. Proceed? [y/n]"
    	read proceed
    	if [ $proceed == "y" ]; then
    		sudo rm /etc/openvswitch/conf.db
    	else 
    		exit 0
    	fi
    fi
    sudo /usr/share/openvswitch/scripts/ovs-ctl --system-id=random start
    sudo ovs-vsctl set Open_vSwitch . other_config:hw-offload="true"
    sudo /usr/share/openvswitch/scripts/ovs-ctl --no-ovsdb-server restart

}

do_stop() {
  	_kubectl delete -f $MANIFESTS_DIR/configMap.yaml || true
  	_kubectl delete -f $MANIFESTS_DIR/configMap-vdpa.yaml || true
  	_kubectl delete -f $MANIFESTS_DIR/sriov_device_plugin.yaml || true
  	_kubectl delete -f $MANIFESTS_DIR/netattach.yaml || true
  	_kubectl delete -f $MANIFESTS_DIR/multus.yaml || true
}

# Main program
CMD=${1:-usage}
case $CMD in
    start)
	[ $# -ne 2 ] && usage
	pfname=$2
        do_start $pfname
        ;;
    stop)
        do_stop $@
        ;;
    usage)
        usage
        ;;
    *)
        usage
        ;;
esac
