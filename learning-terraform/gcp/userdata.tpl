#!/bin/bash
containerd_version=2.0.1
runc_version=1.2.3
cni_plugins_version=1.6.1
calico_version=3.29.1
k8s_version="1.31"
pod_cidr=192.168.0.0/16
nerdctl_version=2.0.2
private_ip=`hostname -I`
# Define URLs
containerd_url="https://github.com/containerd/containerd/releases/download/v$containerd_version/containerd-$containerd_version-linux-amd64.tar.gz"
containerd_service_url="https://raw.githubusercontent.com/containerd/containerd/main/containerd.service"
runc_url="https://github.com/opencontainers/runc/releases/download/v$runc_version/runc.amd64"
cni_plugins_url="https://github.com/containernetworking/plugins/releases/download/v$cni_plugins_version/cni-plugins-linux-amd64-v$cni_plugins_version.tgz"
nerdctl_url="https://github.com/containerd/nerdctl/releases/download/v$nerdctl_version/nerdctl-$nerdctl_version-linux-amd64.tar.gz"
k8s_release_key_url="https://pkgs.k8s.io/core:/stable:/v$k8s_version/deb/Release.key"
tigera_operator_url="https://raw.githubusercontent.com/projectcalico/calico/v$calico_version/manifests/tigera-operator.yaml"
custom_resources_url="https://raw.githubusercontent.com/projectcalico/calico/v$calico_version/manifests/custom-resources.yaml"


# Settings for ipv4, modprobe and iptables
cat <<EOF | tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

modprobe overlay
modprobe br_netfilter

cat <<EOF | tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF

sysctl --system

lsmod | grep br_netfilter
lsmod | grep overlay

sysctl net.bridge.bridge-nf-call-iptables net.bridge.bridge-nf-call-ip6tables net.ipv4.ip_forward

# Prepare container runtime and its dependencies
mkdir -p /opt/downloads
wget $containerd_url -O /opt/downloads/containerd.tar.gz
tar Cxzvf /usr/local /opt/downloads/containerd.tar.gz
mkdir -p /usr/local/lib/systemd/system
wget $containerd_service_url -O /usr/local/lib/systemd/system/containerd.service
systemctl daemon-reload
systemctl enable --now containerd
# install runc
wget $runc_url -O /usr/local/sbin/runc
chmod 0755 /usr/local/sbin/runc
# install cni plugins
mkdir -p /opt/cni/bin
wget $cni_plugins_url -O /opt/downloads/cni-plugins.tgz
tar Cxzvf /opt/cni/bin /opt/downloads/cni-plugins.tgz
# configure containerd and cgroup driver
mkdir -p /etc/containerd
containerd config default > /etc/containerd/config.toml
sed -i 's/SystemdCgroup = false/SystemdCgroup = true/g' /etc/containerd/config.toml
systemctl restart containerd
# Install kubeadm, kubelet and kubectl
apt-get update
apt-get install -y apt-transport-https ca-certificates curl gpg
stat /etc/apt/keyrings
curl -fsSL $k8s_release_key_url | gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v'$k8s_version'/deb/ /' | tee /etc/apt/sources.list.d/kubernetes.list
apt-get update
apt-get install -y kubelet kubeadm kubectl
# Set up k8s cluster with kubeadm
kubeadm init --control-plane-endpoint=$private_ip --apiserver-advertise-address=$private_ip --pod-network-cidr=$pod_cidr
mkdir -p $HOME/.kube
cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
mkdir -p /home/ubuntu/.kube
cp -i /etc/kubernetes/admin.conf /home/ubuntu/.kube/config
chown -R ubuntu. /home/ubuntu/.kube/

# untaint master node
kubectl taint nodes --all node-role.kubernetes.io/control-plane-

# Deploy CNI plugin for pod networking
kubectl create -f $tigera_operator_url
kubectl create -f $custom_resources_url

# Check node status whether ready
kubectl get nodes

# setup nerdctl
wget $nerdctl_url -O /opt/downloads/nerdctl.tar.gz
tar Cxzvvf /usr/local/bin /opt/downloads/nerdctl.tar.gz
mkdir -p /etc/nerdctl/
containerd_unix_sock=`grep -i sock /etc/containerd/config.toml | grep address | awk -F"address = " '{print $2}' | tr -d '"'`
cat <<EOF | tee /etc/nerdctl/nerdctl.toml
# This is an example of /etc/nerdctl/nerdctl.toml .
# Unrelated to the daemon's /etc/containerd/config.toml .

debug          = false
debug_full     = false
address        = "$containerd_unix_sock"
namespace      = "k8s.io"
snapshotter    = "stargz"
cgroup_manager = "cgroupfs"
hosts_dir      = ["/etc/containerd/certs.d", "/etc/docker/certs.d"]
experimental   = true
EOF

# kubeconfig for root user
mkdir -p /root/.kube
cp -i /etc/kubernetes/admin.conf /root/.kube/config