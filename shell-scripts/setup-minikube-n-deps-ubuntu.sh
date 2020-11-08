#!/bin/bash

# For Ubuntu 16.04+

sudo apt-get update
echo "******************************************************************"
echo "Installing docker..."
echo "******************************************************************"
sudo apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common telnet
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io
sleep 3
docker_status=`sudo systemctl is-active docker`
if [[ "$docker_status" != "active" ]]
then
    sudo systemctl start docker
fi
sudo docker run hello-world
echo "Running post installation steps for docker."
sudo usermod -aG docker `whoami`
echo "******************************************************************"
echo "Logout and login again to check if you can run docker without sudo"
echo "******************************************************************"
echo "Installing minikube now..."
echo "******************************************************************"
curl -LO https://storage.googleapis.com/minikube/releases/v1.13.1/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
echo "******************************************************************"
echo "Installing kubectl now..."
echo "******************************************************************"
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.19.2/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
kubectl version

