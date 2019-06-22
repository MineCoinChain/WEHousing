## 	ubuntu 环境需求

```shell
	uname -a
	ls -l /sys/class/misc/device-mapper
```

## 	部署Docker

```shell
//安装基本软件
	apt-get update
	apt-get install apt-transport-https ca-certificates curl software-properties-common -y    
```

```shell
//使用官方推荐源{不推荐}
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
     add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"    
```

```shell
//使用阿里云的源{推荐}
    curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
     add-apt-repository "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
```

```shell
//软件源升级
 apt-get update 
```

```shell
//安装docker
 apt-get install docker-ce -y   
 //可以指定版本安装docker：
 apt-get install docker-ce=<VERSION> -y
```

```shell
//查看支持的docker版本
apt-cache madison docker-ce 
```

```shell
//测试docker
docker version
```

```shell
docker加速器配置
	curl -sSL https://get.daocloud.io/daotools/set_mirror.sh | sh -s http://e5d212cc.m.daocloud.io
//修改daemon.json文件，增加后边的内容
{"registry-mirrors": ["http://e5d212cc.m.daocloud.io"], "insecure-registries": []}
//注意：
//docker cloud加速器的默认内容是少了一条配置，所以我们要编辑文件在后面加上绿色背景的内容，然后再重启docker
```

```docker
//重启docker
systemctl restart docker
//拉取镜像测试
docker pull mysql
docker pull nginx
docker pull ubuntu
```

​	**到这里docker已经安装成功了**

## 删除安装包的命令

```shell
删除docker命令：
:~$  apt-get purge docker-ce -y
:~$  rm -rf /etc/docker                //docker的认证目录
:~$  rm -rf /var/lib/docker/            //docker的应用目录
```

## 解决docker权限问题

```
docker权限问题
#方法1：一劳永逸
	#如果还没有 docker group 就添加一个：
	:~$sudo groupadd docker
	#将用户加入该 group 内。然后退出并重新登录就生效啦。
	:~$sudo gpasswd -a ${USER} docker
	#重启 docker 服务
	:~$systemctl restart docker
	#切换当前会话到新 group 或者重启 X 会话
	:~$newgrp - docker
	#注意:最后一步是必须的，否则因为 groups 命令获取到的是缓存的组信息，刚添加的组信息未能生效，
	所以 docker images 执行时同样有错。
#方法2：
	#每次启动docker或者重启docker的之后
	:~$cd /var/run
	:~$sudo chmod 666 docker.sock
	#方法3：每条命令前面加上sudo
```

