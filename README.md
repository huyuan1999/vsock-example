**版本要求**
* 宿主机内核版本 4.8+
* 虚拟机内核版本 4.8+
* qemu-kvm版本 2.8+

**当前环境**
**hypervisor kernel**
* Linux localhost 4.15.0-159-generic #167-Ubuntu SMP Tue Sep 21 08:55:05 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux

**VM kernel**
* Linux vhost 4.15.0-159-generic #167-Ubuntu SMP Tue Sep 21 08:55:05 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux

**qemu-system-x86_64**
* QEMU emulator version 2.11.1(Debian 1:2.11+dfsg-1ubuntu7.38)

#### 环境配置
```bash
# 安装 qemu 和 virt
$ apt -y install qemu libvirt-bin libvirt-dev virtinst

# 设置 libvirtd 开机自启动并 start libvirtd
$ systemctl enable --now libvirtd

# 卸载 hypervisor 上相关内核模块
$ modprobe -r vmw_vsock_vmci_transport
$ modprobe -r vmw_vsock_virtio_transport_common
$ modprobe -r vsock

# 加载 hypervisor 上的 vhost_vsock 模块
$ modprobe vhost_vsock

# 安装虚拟机
$ virt-install \
-n kvm-01 \
--memory 16384 \
--vcpus 4 \
--cdrom=/opt/ubuntu-18.04.5-live-server-amd64.iso \
--disk /opt/kvm/kvm-01.qcow2,size=50 \
--graphics vnc,listen=0.0.0.0,port=6001 \
--os-type=linux

# 停止虚拟机
$ virsh shutdown kvm-01

# 把这一段放到脚本里面, 扔到后台执行, 要不然会断网。或者通过配置文件创建网桥并关联网卡
$ ip link add br0 type bridge
$ ip link set br0 up
$ ip addr add 10.10.1.135/24 dev br0
$ ip addr del 10.10.1.135/24 dev eno2
$ ip link set eno2 master br0
$ ip route del default
$ ip route add default via 10.10.1.1

# 启动虚拟机插入 vsock 设备
$ qemu-system-x86_64 -m 16G -hda /opt/kvm/kvm-01.qcow2 \
-device vhost-vsock-pci,id=vhost-vsock-pci0,guest-cid=3 \
-vnc :0 --enable-kvm -daemonize \
-net nic \
-net tap,ifname=vnet01,script=/etc/qemu-ifup,downscript=qemu-ifdown
```



