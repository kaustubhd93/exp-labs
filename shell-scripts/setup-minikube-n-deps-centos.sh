#!/bin/bash
echo "Setting up minikube..."
sudo yum install -y telnet vim git
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
minikube version
echo "Setting up kubectl..."
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.19.2/bin/linux/amd64/kubectl
sudo chmod +x kubectl
sudo mv kubectl /usr/local/bin/
kubectl version
echo "Setting up docker..."
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install docker-ce docker-ce-cli containerd.io
docker_status=`sudo systemctl is-active docker`
if [[ "$docker_status" != "active" ]]
then
    sudo systemctl start docker
fi
sudo docker run hello-world
sudo usermod -aG docker `whoami`
echo "******************************************************************"
echo "Logout and login again to check if you can run docker without sudo"
echo "******************************************************************"

