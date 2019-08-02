# Go+Micro+Fabric微服务+区块链项目实战-WeHousing房屋短租+房产认证平台

微服务实现的基本功能:用户注册，用户登录，头像上传，用户详细信息获取，实名认证检测，房源发布，首页展示，搜索房源，订单管理，用户评价等服务。同时我们通过fabric溯源技术对房产信息进行认证，告别二手房东的烦恼！

__PS:__
     __1.可作为微服务和联盟链fabric的入门教程，一个项目玩转目前golang最火的两个领域：微服务和区块链。__
     __2.前端页面稍微有点儿low，不过完全可以换个前端做出高大上的效果。__
     __3.如果对区块链技术不太感冒也不用担心，在实现上我采用了微服务和区块链分离的策略，首先通过微服务实现项目的基本功能，随后通过fabric对房产信息进行认证存储。只想学习微服务的同学可以绕过fabric这一部分。__
     __4.fabric环境搭建比较复杂，请克制砸电脑的冲动。__
     __5.我的仓库里有一个HouseTracing的项目，没事也可以看看，这是一个基于fabric的房产溯源认证平台。__


## 技术栈
+ golang + docker + consul + grpc + protobuf + beego + mysql + redis + fastDFS + nginx + fabric


## 目标功能
- [x] 功能模块
    - [x] 用户模块
        - [x] 注册
          - [x] 获取验证码图片服务
          - [x] 获取短信验证码服务
          - [x] 发送注册信息服务
        - [x] 登录
          - [x] 获取session信息服务
          - [x] 获取登录信息服务
        - [x] 退出
        - [x] 个人信息获取
          - [x] 获取用户基本信息服务
          - [x] 更新用户名服务
          - [x] 发送上传用户头像服务
        - [x] 实名认证
          - [x] 获取用户实名信息服务
          - [x] 发送用户实名认证信息服务
    - [x] 房屋模块
        - [x] 首页展示
          - [x] 获取首页轮播图服务
        - [x] 房屋详情
           - [x] 发布房屋详细信息的服务
           - [x] 上传房屋图片的服务
        - [x] 地区列表
        - [x] 房屋搜索
    - [x] 订单模块
        - [x] 订单确认
        - [x] 发布订单
        - [x] 查看订单信息
        - [x] 订单评论

## 项目文档

​	document文件夹下：

1. ​	整体架构图
2. ​	微服务框架图
3. ​	接口文档

## 运行环境配置及使用教程

[docker安装教程](./configurationFile/DockerInstall.md)

[protobuf安装及使用教程](./configurationFile/protobuf.md)

[micro介绍及安装教程](./configurationFile/micro.md)

[FastDFS+nginx环境配置](./configurationFile/nginxAndFastDFS-nginx-moduleDownload.md)

[Consul安装及使用教程](./configurationFile/Consul.md)

[redis安装教程](./configurationFile/redisDownload.md)

## 项目启动

- 项目启动：  
    - **注意: 项目启动前请先查看项目配置环境文件,配置你相应的设置,并安装好各个环境,mysql+redis+nginx+fastDFS+consul+Micro等**

- consul启动：  
    ```    shell
    开发测试过程中可以使用单机模式
    consul agent -dev
    ```

- redis服务端启动  

    ```shell
     sudo redis-server /etc/redis/redis.conf
    ```

- FastDFS服务启动

    ```shell
    sudo fdfs_trackerd /etc/fdfs/tracker.conf
    sudo fdfs_storaged /etc/fdfs/storage.conf
    ```

- nginx
    ```shell
    启动nginx
    sudo /usr/local/nginx/sbin/nginx
    重启nginx
    sudo /usr/local/nginx/sbin/nginx -s reload
    ```
    
- 创建服务的指令
    ```shell
    新创建微服务的指令为,ProjectPath为项目所在路径，ServerName是服务名字
    micro new --type srv ProjectPath/ServerName
    ```
    
- 初始化protobuf文件的命令
    ```shell
    可以直接执行generateproto.sh文件，也可以cd到该项目的项目目录下然后执行如下命令：
    protoc --proto_path=. --go_out=. --micro_out=. proto/example/example.proto
    ```
    
## 项目布局
```
├── DeleteSession
│   ├── 退出登录时清除session
├── GetArea
│   ├── 获取地区信息服务
├── GetImageCd
│   ├── 获取验证码图片服务
├── GetSession
│   ├── 获取Session信息服务
├── GetSmscd
│   ├── 获取短信信息服务
├── GetUserHouses
│   ├── 获取用户已发布房屋的服务
├── GetUserInfo
│   ├── 获取用户详细信息的服务
├── IhomeWeb
│   ├── conf 项目配置文件
│   │   ├── app.conf
│   │   ├── data.sql
│   │   └── redis.conf
│   ├── handler
│   │   └── handler.go 配置路由
│   ├── html 项目静态文件
│   ├── main.go 主函数
│   ├── model 数据库模型
│   │   └── models.go
│   ├── plugin.go
│   ├── server.sh
│   └── utils 项目中用到的工具函数
│       ├── config.go
│       ├── error.go
│       └── misc.go
├── PostAvatar
│   ├──	发送（上传）用户头像服务
├── PostHouses
│   ├── 发送（发布）房源信息服务
├── PostHousesImage
│   ├── 发送（上传）房屋图片服务
├── PostLogin
│   ├── 发送登录服务消息
├── PostRet
│   ├── 发现注册信息服务
├── PostUserAuth
│   ├── 发送用户实名认证信息服务
├── PutUserInfo
│   ├── 发送用户信息
├── GetUserAuth
│   ├── 获取（检查）用户实名信息服务
├── PostHousesImage
│   ├── 发送（上传）房屋图片服务
├── GetHouseInfo
│   ├── 获取房屋详细信息服务
├── GetIndex
│   ├── 获取首页轮播图片服务
├── GetHouses
│   ├── 获取（搜索）房源服务
├── PostOrders
│   ├── 发送（发布）订单服务
├── GetUserOrder
│   ├── 获取房东/租户订单信息服务
├── PutOrders
│   ├── 更新房东同意/拒绝订单
├── PutComments
│   ├── 更新用户评价订单信息
└── README.md

```

## Docker学习教程
##### [Docker学习教程](./Document/Docker/docker.md)

## GO微服务教程（项目具体实现）

##### [01 项目展示](./Document/material/01项目展示.md)

##### [02 微服务的概念](./Document/material/02微服务的概念.md)

##### [03 微服务与单体式的对比](./Document/material/03微服务与单体式的对比.md)

##### [04 protobuf](./Document/material/04protobuf讲义.md)

##### [05 GRPC](./Document/material/05GRPC.md)

##### [06 服务发现](./Document/material/06服务发现.md)

##### [07 代理均衡](./Document/material/07代理均衡.md)

##### [08 Consul](./Document/material/08Consul.md)

##### [09 micro](./Document/material/09micro.md)

##### [10 租房网](./Document/material/10租房网.md)

##### [11 获取地域信息](./Document/material/11获取地域信息.md)

##### [12 获取验证码图片](./Document/material/12获取验证码图片.md)

##### [13 获取短信验证码](./Document/material/13获取短信验证码.md)

##### [14 注册请求](./Document/material/14注册请求.md)

##### [15 获取session信息](./Document/material/15获取session信息.md)

##### [16 登录请求](./Document/material/16登录请求.md)

##### [17 退出登陆](./Document/material/17退出登陆.md)

##### [18 获取用户信息](./Document/material/18获取用户信息.md)

##### [19 上传用户头像](./Document/material/19上传用户头像.md)

##### [20 更新用户名](./Document/material/20更新用户名.md)

##### [21 检查用户实名认证](./Document/material/21检查用户实名认证.md)

##### [22 更新实名认证信息](./Document/material/22更新实名认证信息.md)

##### [23 其他模块](./Document/material/23业务梳理.md)

##### [24 使用docker-compose进行单机集群启动](./Document/material/24使用docker-compose进行单机集群启动.md)
