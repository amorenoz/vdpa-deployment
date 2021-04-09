#!/bin/bash

## GLOBAL VARS
KUBECTL=${KUBECTL:-$(which kubectl)}

### Common SR-IOV functions

get_pci_addr() {
	local pf=$1
	local vf=${2:-}
	if [ -z $vf ]; then
		echo $(basename $(readlink /sys/class/net/${pf}/device))
	else
		echo $(basename $(readlink /sys/class/net/${pf}/device/virtfn${vf}))
	fi
}

set_eswitch_mode() {
	local pf=$1
	local mode=$2
	sudo devlink dev eswitch set pci/$(get_pci_addr ${pf})  mode ${mode}
#	devlink dev eswitch set pci/$(get_pci_addr ${pf})  encap disable
}

set_tc_offload() {
	local netdev=$1
	sudo ethtool -K $netdev hw-tc-offload on
}

get_representor() {
	local pf=$1
	local vfid=$2

	switchid=$(ip -d link show ${pf} | sed -n 's/.* switchid \([^ ]*\).*/\1/p')
	# FXME: The following code including pf0 is not portable
	portname="pf0vf${vfid}"
	for dev in $(ls -x /sys/class/net); do
		[[ -f /sys/class/net/${dev}/phys_port_name ]] || continue
		[[ -f /sys/class/net/${dev}/phys_switch_id ]] || continue
		phys_port_name=$(cat /sys/class/net/${dev}/phys_port_name 2> /dev/null || true)
		phys_switch_id=$(cat /sys/class/net/${dev}/phys_switch_id 2> /dev/null || true)
		if [ "${phys_port_name}" == "$portname" ] && [ "${phys_switch_id}" == "${switchid}" ]; then
			echo "$dev"
			return
		fi
	done
}

## Common Kubernetes functions

wait_pods_ready() {
  local -r tries=30
  local -r wait_time=10

  local -r wait_message="Waiting for all pods to become ready.."
  local -r error_message="Not all pods were ready after $(($tries*$wait_time)) seconds"

  local -r get_pods='_kubectl get pods --all-namespaces'
  local -r action="_check_all_pods_ready"

  set +x
  trap "set -x" RETURN

  if ! retry "$tries" "$wait_time" "$action" "$wait_message" "$get_pods"; then
    error $error_message
  fi

  echo "all pods are ready"
  return 0
}

_check_all_pods_ready() {
  all_pods_ready_condition=$(_kubectl get pods -A --no-headers -o custom-columns=':.status.conditions[?(@.type == "Ready")].status')
  if [ "$?" -eq 0 ]; then
    pods_not_ready_count=$(grep -cw False <<< "$all_pods_ready_condition")
    if [ "$pods_not_ready_count" -eq 0 ]; then
      return 0
    fi
  fi

  return 1
}

_kubectl() {
    ${KUBECTL} "$@"
}

function wait_for_daemonSet {
  local name=$1
  local namespace=$2
  local required_replicas=$3

  if [[ $namespace != "" ]];then
    namespace="-n $namespace"
  fi

  if (( required_replicas < 0 )); then
      echo "DaemonSet $name ready replicas number is not valid: $required_replicas"
      return 1
  fi

  local -r tries=30
  local -r wait_time=10
  wait_message="Waiting for DaemonSet $name to have $required_replicas ready replicas"
  error_message="DaemonSet $name did not have $required_replicas ready replicas"
  action="_kubectl get daemonset $namespace $name -o jsonpath='{.status.numberReady}' | grep -w $required_replicas"

  if ! retry "$tries" "$wait_time" "$action" "$wait_message";then
    echo $error_message
    return 1
  fi

  return  0
}


# Common Misc functions
function error() {
        echo $@ >&2
        exit 1
}

function retry {
  local -r tries=$1
  local -r wait_time=$2
  local -r action=$3
  local -r wait_message=$4
  local -r waiting_action=${5:-""}

  eval $action
  local return_code=$?
  for i in $(seq $tries); do
    if [[ $return_code -ne 0 ]] ; then
      echo "[$i/$tries] $wait_message"
      eval $waiting_action
      sleep $wait_time
      eval $action
      return_code=$?
    else
      return 0
    fi
  done

  return 1
}

