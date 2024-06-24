# 离线安装docker

1. 以rpm的方式安装docker


[参考文档](https://gist.github.com/ShockwaveNN/2e37d61fa04e19ba814667b05502bc1c)

```bash
sudo rpm -i container-selinux-2.107-1.el7_6.noarch.rpm
sudo rpm -i docker-ce-19.03.5-3.el7.x86_64.rpm docker-ce-cli-19.03.5-3.el7.x86_64.rpm docker-ce-selinux-17.03.3.ce-1.el7.noarch.rpm containerd.io-1.2.6-3.3.el7.x86_64.rpm
sudo systemctl start docker
sudo docker load -i hello-world.docker
sudo docker run hello-world
```


2. 以二进制的方式安装docker


1 docker安装
1.1 下载docker

下载地址：https://download.docker.com/linux/static/stable/

1.查看系统架构：

```bash
[root@localhost es]# arch
x86_64

```
2.在下载页选择对应版本的docker安装包，我的是x86的系统，选择的是：https://download.docker.com/linux/static/stable/x86_64/docker-20.10.0.tgz

3.上传docker-20.10.0.tgz安装包到服务器

1.2 安装
解压安装包，并移动可执行文件至/usr/bin/目录：

```bash
$ tar -xzf docker-20.10.0.tgz
$ mv docker/* /usr/bin/
```

1.3 编辑docker的系统服务文件
vi /usr/lib/systemd/system/docker.service
将下面的内容复制到刚创建的docker.service文件中

```bash
cat <<EOF >>/usr/lib/systemd/system/docker.service

[Unit]
 
Description=Docker Application Container Engine
 
Documentation=https://docs.docker.com
 
After=network-online.target firewalld.service
 
Wants=network-online.target
 
 
 
[Service]
 
Type=notify
 
ExecStart=/usr/bin/dockerd
 
ExecReload=/bin/kill -s HUP $MAINPID
 
LimitNOFILE=infinity
 
LimitNPROC=infinity
 
TimeoutStartSec=0
 
Delegate=yes
 
KillMode=process
 
Restart=on-failure
 
StartLimitBurst=3
 
StartLimitInterval=60s
 
 
 
[Install]
 
WantedBy=multi-user.target
EOF
```



添加可执行权限：

```bash
chmod +x /usr/lib/systemd/system/docker.service

systemctl daemon-reload
```
 
编辑daemon.json