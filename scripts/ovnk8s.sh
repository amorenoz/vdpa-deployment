#!/usr/bin/env bash
set -euxo pipefail

SCRIPT=`realpath -s $0`
SCRIPTPATH=`dirname $SCRIPT`

run_kubectl() {
  local retries=0
  local attempts=10
  while true; do
    if kubectl "$@"; then
      break
    fi

    ((retries += 1))
    if [[ "${retries}" -gt ${attempts} ]]; then
      echo "error: 'kubectl $*' did not succeed, failing"
      exit 1
    fi
    echo "info: waiting for 'kubectl $*' to succeed..."
    sleep 1
  done
}


usage() {
    echo "usage: $0 [start|stop]"
}

parse_args() {
    CMD=$1
}

print_params() {
     echo ""
     echo "OVN_HA = $OVN_HA"
     echo "OVN_GATEWAY_MODE = $OVN_GATEWAY_MODE"
     echo "OVN_HYBRID_OVERLAY_ENABLE = $OVN_HYBRID_OVERLAY_ENABLE"
     echo "OVN_DISABLE_SNAT_MULTIPLE_GWS = $OVN_DISABLE_SNAT_MULTIPLE_GWS"
     echo "OVN_EMPTY_LB_EVENTS = $OVN_EMPTY_LB_EVENTS"
     echo "OVN_MULTICAST_ENABLE = $OVN_MULTICAST_ENABLE"
     echo "OVN_IMAGE = $OVN_IMAGE"
     echo "MASTER_LOG_LEVEL = $MASTER_LOG_LEVEL"
     echo "NODE_LOG_LEVEL = $NODE_LOG_LEVEL"
     echo "DBCHECKER_LOG_LEVEL = $DBCHECKER_LOG_LEVEL"
     echo "OVN_LOG_LEVEL_NORTHD = $OVN_LOG_LEVEL_NORTHD"
     echo "OVN_LOG_LEVEL_NB = $OVN_LOG_LEVEL_NB"
     echo "OVN_LOG_LEVEL_SB = $OVN_LOG_LEVEL_SB"
     echo "OVN_LOG_LEVEL_CONTROLLER = $OVN_LOG_LEVEL_CONTROLLER"
     echo "OVN_LOG_LEVEL_NBCTLD = $OVN_LOG_LEVEL_NBCTLD"
     echo ""
}

set_default_params() {
  # Set default values
  K8S_VERSION=${K8S_VERSION:-v1.20.0}
  OVN_GATEWAY_MODE=${OVN_GATEWAY_MODE:-local}
  INSTALL_INGRESS=${INSTALL_INGRESS:-false}
  OVN_HA=${OVN_HA:-false}
  REMOVE_TAINT=${REMOVE_TAINT:-true}
  OVN_HYBRID_OVERLAY_ENABLE=${OVN_HYBRID_OVERLAY_ENABLE:-false}
  OVN_DISABLE_SNAT_MULTIPLE_GWS=${OVN_DISABLE_SNAT_MULTIPLE_GWS:-false}
  OVN_EMPTY_LB_EVENTS=${OVN_EMPTY_LB_EVENTS:-false}
  OVN_MULTICAST_ENABLE=${OVN_MULTICAST_ENABLE:-false}
  OVN_IMAGE=${OVN_IMAGE:-local}
  MASTER_LOG_LEVEL=${MASTER_LOG_LEVEL:-5}
  NODE_LOG_LEVEL=${NODE_LOG_LEVEL:-5}
  DBCHECKER_LOG_LEVEL=${DBCHECKER_LOG_LEVEL:-5}
  OVN_LOG_LEVEL_NORTHD=${OVN_LOG_LEVEL_NORTHD:-"-vconsole:info -vfile:info"}
  OVN_LOG_LEVEL_NB=${OVN_LOG_LEVEL_NB:-"-vconsole:info -vfile:info"}
  OVN_LOG_LEVEL_SB=${OVN_LOG_LEVEL_SB:-"-vconsole:info -vfile:info"}
  OVN_LOG_LEVEL_CONTROLLER=${OVN_LOG_LEVEL_CONTROLLER:-"-vconsole:info"}
  OVN_LOG_LEVEL_NBCTLD=${OVN_LOG_LEVEL_NBCTLD:-"-vconsole:info"}

  NET_CIDR_IPV4=${NET_CIDR_IPV4:-10.244.0.0/16}
  SVC_CIDR_IPV4=${SVC_CIDR_IPV4:-10.96.0.0/16}
  NET_CIDR=$NET_CIDR_IPV4
  SVC_CIDR=$SVC_CIDR_IPV4

  NUM_MASTER=1
  if [ "$OVN_HA" == true ]; then
    NUM_MASTER=3
  fi
}

detect_apiserver_url() {
  # Detect API_URL used for in-cluster communication
  #
  # Despite OVN run in pod they will only obtain the VIRTUAL apiserver address
  # and since OVN has to provide the connectivity to service
  # it can not be bootstrapped
  #
  # This is the address of the node with the control-plane
  API_URL=$(kubectl config view | grep server | awk -e '{print $2}')
}

delete_kube_proxy() {
  kubectl -n kube-system delete ds kube-proxy || true
  kubectl get nodes
}

delete_num_master() {
  # TODO detect properly
  NUM_MASTER=1
}

create_ovn_kube_manifests() {
  pushd $OVN_ROOT/dist/images
  ./daemonset.sh \
    --image="${OVN_IMAGE}" \
    --net-cidr="${NET_CIDR}" \
    --svc-cidr="${SVC_CIDR}" \
    --gateway-mode="${OVN_GATEWAY_MODE}" \
    --hybrid-enabled="${OVN_HYBRID_OVERLAY_ENABLE}" \
    --disable-snat-multiple-gws="${OVN_DISABLE_SNAT_MULTIPLE_GWS}" \
    --ovn-empty-lb-events="${OVN_EMPTY_LB_EVENTS}" \
    --multicast-enabled="${OVN_MULTICAST_ENABLE}" \
    --k8s-apiserver="${API_URL}" \
    --ovn-master-count="${NUM_MASTER}" \
    --ovn-unprivileged-mode=no \
    --master-loglevel="${MASTER_LOG_LEVEL}" \
    --node-loglevel="${NODE_LOG_LEVEL}" \
    --dbchecker-loglevel="${DBCHECKER_LOG_LEVEL}" \
    --ovn-loglevel-northd="${OVN_LOG_LEVEL_NORTHD}" \
    --ovn-loglevel-nb="${OVN_LOG_LEVEL_NB}" \
    --ovn-loglevel-sb="${OVN_LOG_LEVEL_SB}" \
    --ovn-loglevel-controller="${OVN_LOG_LEVEL_CONTROLLER}" \
    --ovn-loglevel-nbctld="${OVN_LOG_LEVEL_NBCTLD}" \
    --egress-ip-enable=true
  popd
}

install_ovn() {
  pushd ${OVN_ROOT}/dist/yaml
  run_kubectl apply -f k8s.ovn.org_egressfirewalls.yaml
  run_kubectl apply -f k8s.ovn.org_egressips.yaml
  run_kubectl apply -f ovn-setup.yaml
  MASTER_NODES=$(kubectl get nodes -l node-role.kubernetes.io/master -o name | sort | head -n "${NUM_MASTER}")
  # We want OVN HA not Kubernetes HA
  # leverage the kubeadm well-known label node-role.kubernetes.io/master=
  # to choose the nodes where ovn master components will be placed
  for n in $MASTER_NODES; do
    kubectl label "$n" k8s.ovn.org/ovnkube-db=true node-role.kubernetes.io/master="" --overwrite
    if [ "$REMOVE_TAINT" == true ]; then
      # do not error if it fails to remove the taint
      kubectl taint node "$n" node-role.kubernetes.io/master:NoSchedule- || true
    fi
  done

  run_kubectl apply -f ovnkube-db.yaml
  run_kubectl apply -f ovnkube-master.yaml
  run_kubectl apply -f ovnkube-node.yaml
  popd
}

start_ovs() {
	if /usr/share/openvswitch/scripts/ovs-ctl status ; then
		echo "OvS already running"
	else
		if [ -f "/etc/openvswitch/conf.db" ]; then
			echo "Do you want to remove the OpenVSwitch configuration? (y/n)"
			read proceed
			if [ $proceed == "y" ]; then
				sudo rm /etc/openvswitch/conf.db
			fi
		fi
		sudo systemctl start openvswitch
    		sudo ovs-vsctl --no-wait set Open_vSwitch . other_config:hw-offload="true"
		sudo systemctl restart openvswitch
	fi
}

stop_ovs() {
	sudo systemctl stop openvswitch || true
	sudo ovs-dpctl del-dp ovs-system || true
	sudo rmmod openvswitch || true
}

kubectl_wait_pods() {
  OVN_TIMEOUT=300s
  if ! kubectl wait -n ovn-kubernetes --for=condition=ready pods --all --timeout=${OVN_TIMEOUT} ; then
    echo "some pods in OVN Kubernetes are not running"
    kubectl get pods -A -o wide || true
    exit 1
  fi
  # core-dns pods might be on a crashloop. Restart them to make the waiting time shorter
  for pod in $(kubectl get pods -n kube-system -l k8s-app=kube-dns -o name); do 
	  kubectl -n kube-system delete $pod
  done

  if ! kubectl wait -n kube-system --for=condition=ready pods --all --timeout=300s ; then
    echo "some pods in the system are not running"
    kubectl get pods -A -o wide || true
    exit 1
  fi
}

sleep_until_pods_settle() {
  echo "Pods are all up, allowing things settle for 30 seconds..."
  sleep 30
}

do_start() {
  set_default_params
  print_params
  detect_apiserver_url
  create_ovn_kube_manifests
  start_ovs
  install_ovn
  delete_kube_proxy
  kubectl_wait_pods
  sleep_until_pods_settle
}

do_stop() {
  pushd ${OVN_ROOT}/dist/yaml
  	kubectl delete -f ovnkube-node.yaml --wait || true
  	kubectl delete -f ovnkube-master.yaml --wait || true
    	kubectl delete -f ovnkube-db.yaml --wait || true
	kubectl apply -f ${SCRIPTPATH}/cleaner.yaml || true
	sleep 10
	kubectl get pods -A
	sleep 5
	ip -4 a
	ip link
	echo "How are things?"
	read waitamin
	## Run cleaner
	## Wait for completion
	## Delete cleaner
  	kubectl delete -f ovn-setup.yaml --wait || kubectl delete namespace ovn-kubernetes --wait || true 
  popd
  #kubectl delete namespace ovn-kubernetes || true
  stop_ovs
  sleep 10
}

CMD=${1:-usage}
case $CMD in
	start)
		do_start	
		exit 0
		;;
	stop)
		do_stop
		exit 0
		;;
	usage)
		usage
		exit 0
		;;
	*)
		usage
		exit 1
		;;
esac
