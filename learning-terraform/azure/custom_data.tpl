#!/bin/bash
sudo apt-get update
sudo apt-get install -y zip unzip
private_ip=`hostname -I`
echo "********************************************************************"
echo $private_ip
echo "********************************************************************"

