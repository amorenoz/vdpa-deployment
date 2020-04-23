# To build:
#  docker build --rm -t vdpa-daemonset .
#

# -------- Builder stage.
FROM centos:7
MAINTAINER Billy McFall <bmcfall@redhat.com>

#
# Install required packages
#
RUN rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO && curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo
RUN yum groupinstall -y "Development Tools"
RUN yum install -y wget numactl-devel git golang make libnl3-devel libudev-devel cmake; yum clean all
# Debug Tools (if needed):
#RUN yum install -y pciutils iproute; yum clean all

#
# Download and Build rmda-core
#
WORKDIR /usr/src
ENV RDMA_CORE_VER 25.4
ENV RDMA_CORE_PATH=/usr/src/rdma-core-$RDMA_CORE_VER
RUN wget https://github.com/linux-rdma/rdma-core/releases/download/v${RDMA_CORE_VER}/rdma-core-${RDMA_CORE_VER}.tar.gz
RUN tar xvzf rdma-core-${RDMA_CORE_VER}.tar.gz
WORKDIR ${RDMA_CORE_PATH}
RUN ./build.sh

#
# Download and Build DPDK
# 
# Stable:
#ENV DPDK_VER 19.08
#ENV DPDK_DIR /usr/src/dpdk-${DPDK_VER}
#WORKDIR /usr/src/
#RUN wget http://fast.dpdk.org/rel/dpdk-${DPDK_VER}.tar.xz
#RUN tar -xpvf dpdk-${DPDK_VER}.tar.xz

# Master:
ENV DPDK_DIR /usr/src/dpdk
WORKDIR /usr/src/
RUN git clone https://github.com/DPDK/dpdk.git dpdk

WORKDIR ${DPDK_DIR}
RUN git checkout master

ENV RTE_TARGET=x86_64-native-linuxapp-gcc
ENV RTE_SDK=${DPDK_DIR}
WORKDIR ${DPDK_DIR}

# Lock down code to commit on 11/13/2019
RUN git checkout 32bf71c3b6774faa0029a62bf6ee9334a7980c84

# DPDK_VER 19.08
RUN sed -i -e 's/EAL_IGB_UIO=y/EAL_IGB_UIO=n/' config/common_linux
RUN sed -i -e 's/KNI_KMOD=y/KNI_KMOD=n/' config/common_linux
RUN sed -i -e 's/LIBRTE_KNI=y/LIBRTE_KNI=n/' config/common_linux
RUN sed -i -e 's/LIBRTE_PMD_KNI=y/LIBRTE_PMD_KNI=n/' config/common_linux
RUN sed -i -e 's/LIBRTE_PMD_KNI=y/LIBRTE_PMD_KNI=n/' config/common_linux
RUN sed -i -e 's/CONFIG_RTE_LIBRTE_MLX5_PMD=n/CONFIG_RTE_LIBRTE_MLX5_PMD=y/' config/common_base
# Additional Debug if Needed
#RUN sed -i -e 's/CONFIG_RTE_LIBRTE_ETHDEV_DEBUG=n/CONFIG_RTE_LIBRTE_ETHDEV_DEBUG=y/' config/common_base

# Add vhost patch
# NOTE: Already in origin/kubecon_vdpa_drivers
COPY ./vhost_substitute.sh ./vhost_substitute.sh
RUN ./vhost_substitute.sh

# Temp patch to allow vDPA Sample APP to work with
# Cascade Glacier PF. Used to simulate Alibaba deployment.
#COPY ./vdpa-dpdk-image/dpdk_ifc_tmp_0001.patch ./dpdk_ifc_tmp_0001.patch
#RUN patch -p1 < dpdk_ifc_tmp_0001.patch

# Add extra debug flags if needed.  EXTRA_CFLAGS="-O0 -g " 
RUN make install -j16 T=${RTE_TARGET} DESTDIR=${RTE_SDK} EXTRA_CFLAGS="-I${RDMA_CORE_PATH}/build/include" EXTRA_LDFLAGS="-L${RDMA_CORE_PATH}/build/lib"

#
# Build userspace-vdpa
#
ENV VDPAD_DIR=/usr/src/userspace-vdpa/uvdpad
WORKDIR /usr/src
RUN git clone https://gitlab.com/mcoquelin/userspace-vdpa.git
WORKDIR ${VDPAD_DIR}
# Add extra debug flags if needed.  EXTRA_CFLAGS="-O0 -g " 
RUN yum install -y jansson-devel
# Patch uvdpad
ADD 0001-Change-socket-path.patch .
RUN git apply 0001-Change-socket-path.patch
RUN make  EXTRA_CFLAGS="-I${RDMA_CORE_PATH}/build/include" EXTRA_LDFLAGS="-L${RDMA_CORE_PATH}/build/lib"
RUN cp build/uvdpad /usr/bin/uvdpad

# TODO: Move just after rdma build
# PseudoInstall: Make library files accessible from a generic directory so that
# they can be copied in the next build-stage
RUN ln -s ${RDMA_CORE_PATH}/build/lib/ /usr/lib64/rdma-core
#

# -------- Import stage.
# BEGIN - Docker 17.05 or higher
FROM centos
COPY --from=0 /usr/bin/uvdpad /usr/bin/uvdpad
COPY --from=0 /usr/lib64/libnuma.so.1 /usr/lib64/libnuma.so.1
COPY --from=0 /usr/lib64/libnl-route-3.so.200 /usr/lib64/libnl-route-3.so.200
COPY --from=0 /usr/lib64/libnl-3.so.200 /usr/lib64/libnl-3.so.200
COPY --from=0 /usr/lib64/libjansson.so.4 /usr/lib64/libjansson.so.4
COPY --from=0 /usr/lib64/rdma-core/. /usr/lib64/.

# END - Docker 17.05 or higher

LABEL io.k8s.display-name="vDPA Daemonset"

ADD ./entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]

