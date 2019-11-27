# To build:
#  docker build --rm -t vdpa-daemonset .
#

# -------- Builder stage.
FROM fedora:29
MAINTAINER Adrian Moreno <amorenoz@redhat.com>

#
# Install required packages
#
RUN yum groupinstall -y "Development Tools"
RUN yum install -y wget numactl-devel git make

#
# Download and Build rmda-core
#
RUN yum -y install cmake gcc libnl3-devel libudev-devel pkgconfig valgrind-devel ninja-build python3-devel python3-Cython; yum clean all
WORKDIR /usr/src
ENV RDMA_CORE_VER stable-v25
ENV RDMA_CORE_PATH=/usr/src/rdma-core-$RDMA_CORE_VER
WORKDIR /usr/src/
RUN wget https://github.com/linux-rdma/rdma-core/archive/${RDMA_CORE_VER}.zip
RUN ls; unzip  ${RDMA_CORE_VER}.zip
WORKDIR ${RDMA_CORE_PATH}
RUN ls -lha 
RUN ./build.sh

#
# Download and Build DPDK
#
ENV DPDK_DIR /usr/src/vdpaDriver/
WORKDIR /usr/src/
RUN wget http://file.mad.redhat.com/~amorenoz/vdpaDriver.tar.gz
RUN tar -xpvf vdpaDriver.tar.gz

WORKDIR ${DPDK_DIR}
ADD vdpa-dpdk-image/mlx/0002-mlx5-fix-debug-compilation.patch ${DPDK_DIR}
ADD vdpa-dpdk-image/mlx/0001-vhost-configure-vDPA-as-soon-as-the-device-is-ready.patch ${DPDK_DIR}
ADD vdpa-dpdk-image/mlx/0002-vhost-add-support-for-memory-hotplug-with-vDPA-devic.patch ${DPDK_DIR}
RUN patch -p1 < 0002-mlx5-fix-debug-compilation.patch
RUN patch -p1 < 0001-vhost-configure-vDPA-as-soon-as-the-device-is-ready.patch
RUN patch -p1 < 0002-vhost-add-support-for-memory-hotplug-with-vDPA-devic.patch

ENV RTE_TARGET=x86_64-native-linuxapp-gcc
ENV RTE_SDK=${DPDK_DIR}

# DPDK_VER 19.08
RUN sed -i -e 's/EAL_IGB_UIO=y/EAL_IGB_UIO=n/' config/common_linux
RUN sed -i -e 's/KNI_KMOD=y/KNI_KMOD=n/' config/common_linux
RUN sed -i -e 's/LIBRTE_KNI=y/LIBRTE_KNI=n/' config/common_linux
RUN sed -i -e 's/LIBRTE_PMD_KNI=y/LIBRTE_PMD_KNI=n/' config/common_linux
RUN sed -i -e 's/CONFIG_RTE_LIBRTE_MLX5_PMD=n/CONFIG_RTE_LIBRTE_MLX5_PMD=y/' config/common_base
RUN sed -i -e 's/CONFIG_RTE_LIBRTE_MLX5_DEBUG=n/CONFIG_RTE_LIBRTE_MLX5_DEBUG=y/' config/common_base

# Additional Debug if Needed
#RUN sed -i -e 's/CONFIG_RTE_LIBRTE_ETHDEV_DEBUG=n/CONFIG_RTE_LIBRTE_ETHDEV_DEBUG=y/' config/common_base


RUN make -j4 install T=${RTE_TARGET} DESTDIR=${RTE_SDK} EXTRA_CFLAGS="-g -I$RDMA_CORE_PATH/build/include " EXTRA_LDFLAGS="-L$RDMA_CORE_PATH/build/lib"

#
# Build vdpa
#
WORKDIR ${DPDK_DIR}/examples/vdpa
COPY ./vdpa-dpdk-image/main_write_table.txt ./main_write_table.txt
COPY ./vdpa-dpdk-image/main_substitute.sh ./main_substitute.sh
RUN ./main_substitute.sh
RUN make EXTRA_CFLAGS="-I$RDMA_CORE_PATH/build/include " EXTRA_LDFLAGS="-L$RDMA_CORE_PATH/build/lib"
RUN cp build/vdpa /usr/bin/vdpa


# -------- Import stage.
# BEGIN - Docker 17.05 or higher
FROM centos
COPY --from=0 /usr/bin/vdpa /usr/bin/vdpa
COPY --from=0 /usr/lib64/libnuma.so.1 /usr/lib64/libnuma.so.1
COPY --from=0 /usr/lib64/libnl-3.so.200 /usr/lib64/libnl-3.so.200
COPY --from=0 /usr/lib64/libnl-route-3.so.200 /usr/lib64/libnl-route-3.so.200
COPY --from=0 /usr/src/rdma-core-stable-v25/build/lib/libibverbs.so.1 /usr/lib64/libibverbs.so.1
COPY --from=0 /usr/src/rdma-core-stable-v25/build/lib/libmlx4.so.1 /usr/lib64/libmlx4.so.1
COPY --from=0 /usr/src/rdma-core-stable-v25/build/lib/libmlx5.so.1 /usr/lib64/libmlx5.so.1
# END - Docker 17.05 or higher

LABEL io.k8s.display-name="vDPA Daemonset"

RUN yum install -y epel-release; yum install -y inotify-tools; yum clean all

ADD ./vdpa-dpdk-image/entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]

