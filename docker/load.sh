#!/usr/bin/env bash
wget --content-disposition https://nct.onlyoffice.com/sh/Xsy  -P ./docker # Hello world image
wget http://mirror.centos.org/centos/7/extras/x86_64/Packages/container-selinux-2.107-1.el7_6.noarch.rpm -P ./docker
wget https://download.docker.com/linux/centos/7/x86_64/stable/Packages/docker-ce-19.03.5-3.el7.x86_64.rpm -P ./docker
wget https://download.docker.com/linux/centos/7/x86_64/stable/Packages/docker-ce-cli-19.03.5-3.el7.x86_64.rpm -P ./docker
wget https://download.docker.com/linux/centos/7/x86_64/stable/Packages/docker-ce-selinux-17.03.3.ce-1.el7.noarch.rpm -P ./docker
wget https://download.docker.com/linux/centos/7/x86_64/stable/Packages/containerd.io-1.2.6-3.3.el7.x86_64.rpm -P ./docker

tar -zcvf docker-1.17-Linux-x86_64.tar.gz docker