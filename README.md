# Go+Micro微服务实战项目-租房网项目

功能:用户注册，用户登录，头像上传，用户详细信息获取，实名认证检测，房源发布，首页展示，搜索房源，订单管理，用户评价等模块。

__注:个人学习微服务使用__

## 技术栈
+ golang + docker + consul + grpc + protobuf + beego + mysql + redis + fastDFS + nginx 


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
└── README.md

```

## 项目Docker部署

​	

## GO微服务教程（项目具体实现）

##### 01 项目展示

##### 02 微服务的概念

##### 03 微服务与单体式的对比

##### 04 protobuf

##### 05 GRPC

##### 06 服务发现

##### 07 代理均衡

##### 08 Consul

##### 09 micro

##### 10 租房网

##### 11 获取地域信息

##### 12 获取验证码图片

##### 13 获取短信验证码

##### 14 注册请求

##### 15 获取session信息

##### 16 登录请求

##### 17 退出登陆

##### 18 获取用户信息

##### 19 上传用户头像

##### 20更新用户名

##### 21检查用户实名认证

##### 22更新实名认证信息

##### 23业务梳理

##### 24使用docker-compose进行单机集群启动
