# 使用docker-compose进行单机集群启动

## docker的安装

```shell
#安装基本软件
$apt-get update
$apt-get install apt-transport-https ca-certificates curl software-properties-common lrzsz -y
#使用阿里云的源{推荐}
$ sudo curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
$ sudo add-apt-repository "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
#软件源升级
$ sudo apt-get update

#安装docker
$ sudo apt-get install docker-ce -y
#测试docker
docker version
#加速器配置
$ curl -sSL https://get.daocloud.io/daotools/set_mirror.sh | sh -s http://f1361db2.m.daocloud.io
#修改配置文件
$ sudo vim /etc/docker/daemon.json
#文件内容
{"registry-mirrors": ["http://f1361db2.m.daocloud.io"], "insecure-registries": []}

#修改权限
#如果还没有 docker group 就添加一个：
$sudo groupadd docker
#将用户加入该 group 内。然后退出并重新登录就生效啦。
$sudo gpasswd -a ${USER} docker
#重启 docker 服务
$systemctl restart docker
#切换当前会话到新 group 或者重启 X 会话
$newgrp - docker
#注意:最后一步是必须的，否则因为 groups 命令获取到的是缓存的组信息，刚添加的组信息未能生效，
#所以 docker images 执行时同样有错。
```



## docker-compose的安装

```shell
#安装依赖工具
sudo apt-get install python-pip -y
#安装编排工具
sudo pip install docker-compose
#查看编排工具版本
sudo docker-compose version
#查看命令帮助
docker-compose --help
#用pip安装依赖包时默认访问https://pypi.python.org/simple/，
#但是经常出现不稳定以及访问速度非常慢的情况，国内厂商提供的pipy镜像目前可用的有：

#在当前用户目录下创建.pip文件夹
mkdir ~/.pip
#然后在该目录下创建pip.conf文件填写：
[global]
trusted-host=mirrors.aliyun.com
index-url=http://mirrors.aliyun.com/pypi/simple/
```



## 部署前的项目修改

### 代码的修改

将所有读取conf文件夹内容部的程序的绝对路径修改为相对路径

```go
//fastdfs中的读取client.conf
fdfs_client.NewFdfsClient("./conf/client.conf")
//utils中读取app.conf
config.NewConfig("ini", "./conf/app.conf")

```

将conf文件复制到各个微服务项目文件夹中

```shell
#仅将所需要的配置文件进行拷贝就可以了
.
├── app.conf	#项目配置信息
└── client.conf #fastdfs客户端配置信息
```





## 项目的编译



```shell
#二进制编译
$ CGO_ENABLED=0 GOOS=linux  /usr/local/go/bin/go build -a -installsuffix cgo -ldflags '-w' -i -o homeweb-web main.go
#编译需要在root账户下进行
#指明cgo工具是否可用的标识在这里表示禁用
CGO_ENABLED=0 
#目标平台（编译后的目标平台）的操作系统（darwin、freebsd、linux、windows）
GOOS=linux  
#由于没有在root下安装go所以我们需要使用go的绝对路径进行使用
/usr/local/go/bin/go build 
#强制重新编译所有涉及的go语言代码包
-a 
#为了使当前的输出目录与默认的编译输出目录分离，可以使用这个标记。此标记的值会作为结果文件的父目录名称的后缀。
-installsuffix 

cgo
# 给 cgo指定命令 
-ldflags 
#关闭所有警告信息
'-w'
#标志安装目标的依赖包。
-i 
#命名
-o ihomeweb 
#编译的main.go地址
./main.go
```



## 服务容器化

web

```Dockerfile
FROM alpine:3.2
#拷贝文件
ADD conf /conf
#拷贝文件
ADD html /html
#拷贝二进制
ADD ihomeweb /ihomeweb
WORKDIR /

ENTRYPOINT [ "/ihomeweb" ]

EXPOSE 8999
```

srv

```
FROM alpine:3.2
ADD conf /conf

ADD getarea-srv /getarea-srv
ENTRYPOINT [ "/getarea-srv" ]
```



## Compose编排

```yaml
consul:
  #覆盖启动后的执行命令
  command: agent -server -bootstrap-expect=1  -node=node1 -client 0.0.0.0 -ui -bind=0.0.0.0 -join 127.0.0.2
  #command: agent -server -bootstrap -rejoin -ui
  #镜像：镜像名称:版本号
  image: consul:latest
  #主机名
  hostname: "registry"
  #暴露端口
  ports:
  - "8300:8300"
  - "8400:8400"
  - "8500:8500"
  - "8600:53/udp"

#web主页
web:
  #覆盖启动后的执行命令
  command: --registry_address=registry:8500 --register_interval=5 --register_ttl=10 web
  #镜像构建的dockerfile文件地址
  build: ./ihomeweb
  links:
  - consul
  ports:
  - "8999:8999"
#获取地区
getarea:
  #覆盖启动后的执行命令
  command: --registry_address=registry:8500 --register_interval=5 --register_ttl=10 srv
  #镜像构建的dockerfile文件地址
  build: ./getarea
  links:
  - consul

#注册三部曲
getimagecd:
  #覆盖启动后的执行命令
  command: --registry_address=registry:8500 --register_interval=5 --register_ttl=10 srv
  #镜像构建的dockerfile文件地址
  build: ./getimagecd
  links:
  - consul

getsmscd:
  #覆盖启动后的执行命令
  command: --registry_address=registry:8500 --register_interval=5 --register_ttl=10 srv
  #镜像构建的dockerfile文件地址
  build: ./getsmscd
  links:
  - consul

postret:
  #覆盖启动后的执行命令
  command: --registry_address=registry:8500 --register_interval=5 --register_ttl=10 srv
  #镜像构建的dockerfile文件地址
  build: ./postret
  links:
  - consul

```





```shell
docker-compose scale getarea=2 
```



