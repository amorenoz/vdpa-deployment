#!/bin/sh

set -e

VDPA_SYS_BINARY_DIR="/usr/bin"
VDPA_VHOST_SOCKETDIR="/var/run/uvdpa"
CLI_PARAMS=""

# Remove any existing vhost socketfiles then recreated the directory
rm -rf $VDPA_VHOST_SOCKETDIR
mkdir -p $VDPA_VHOST_SOCKETDIR

$VDPA_SYS_BINARY_DIR/uvdpad $CLI_PARAMS
