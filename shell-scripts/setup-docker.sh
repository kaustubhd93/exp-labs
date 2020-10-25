#!/bin/bash

# For Ubuntu 16.04+

sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io
sleep 3
sudo systemctl status docker
sudo docker run hello-world
echo "Post installation steps for docker."
sudo usermod -aG docker `whoami`
echo "******************************************************************"
echo "Logout and login again to check if you can run docker without sudo"
echo "******************************************************************"

