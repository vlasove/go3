#!/bin/bash

sudo apt-get update
sudo apt-get -y upgrade

curl -O https://dl.google.com/go/go1.13.7.linux-amd64.tar.gz

tar -xvf go1.13.7.linux-amd64.tar.gz

sudo mv go /usr/local