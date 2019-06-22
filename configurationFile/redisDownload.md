# redis
### redis下载
1. 当前下载的版本为稳定版3.2版本  
    `sudo wget http://download.redis.io/releases/redis-3.2.8.tar.gz`
2. 解压  
    `tar -zxvf redis-3.2.8.tar.gz`
3. 放到/usr/local目录下  
    `sudo mv ./redissudo mv ./redis`  
4. 进入redis目录  
    `cd /usr/local/redis/`
5. 生成  
    `sudo make`
6. 测试,这段运行时间会较长  
    `sudo make intall`
7. 安装,将redis的命令安装到/usr/local/bin/⽬录  
    `sudo make install`
8. 安装完成后，我们进入目录/usr/local/bin中查看  
    `cd /usr/local/bin`  
    `ls -all`  
    ![image](./images/p1_12.png)
    > redis-server redis服务器  
    redis-cli redis命令行客户端  
    redis-benchmark redis性能测试工具  
    redis-check-aof AOF文件修复工具  
    redis-check-rdb RDB文件检索工具  
9. 配置⽂件，移动到/etc/⽬录下  
   配置⽂件⽬录为/usr/local/redis/redis.conf  
   > sudo cp /usr/local/redis/redis.conf /etc/redis/
---
### 下载过程中可能出现的错误
   - 安装redis时遇到You need tcl 8.5 or newer in order to run the Redis test  
   下载安装tcl：  
    1. sudo wget http://downloads.sourceforge.net/tcl/tcl8.6.8-src.tar.gz   
    2. sudo tar xzvf tcl8.6.8-src.tar.gz （解压到usr/local/下的)  
    3. cd  /usr/local/tcl8.6.8/unix/  
    4. sudo ./configure  
    5. sudo make  
    6. sudo make install  
    - 由于tcl下载的服务器在外国,可能下载不了,可以使用我翻墙下载好的包[tcl8.6.1-src.tar.gz](https://pan.baidu.com/s/1NkK7VbeNBrbTPUeTxcYD6A),然后继续执行2-6步 
### redis配置
   - redis配置信息在/etc/redis/redis.conf下  
   `sudo vi /etc/redis/redis.conf`
   
   - 绑定ip：如果需要远程访问，可将此⾏注释，或绑定⼀个真实ip  
    `bind 127.0.0.1`

   - 端⼝，默认为6379  
    `port 6379`

   - 是否以守护进程运⾏  
    - 如果以守护进程运⾏，则不会在命令⾏阻塞，类似于服务
    如果以⾮守护进程运⾏，则当前终端被阻塞
    设置为yes表示守护进程，设置为no表示⾮守护进程
    推荐设置为yes  
    `daemonize yes`

   - 数据⽂件  
    `dbfilename dump.rdb`

   - 数据⽂件存储路径  
    `dir /var/lib/redis`

   - ⽇志⽂件  
    `logfile /var/log/redis/redis-server.log`

   - 数据库，默认有16个  
    `database 16`

   - 主从复制，类似于双机备份  
    `slaveof`
   - redis配置信息参考信息  
   [redis配置信息参考](http://blog.csdn.net/ljphilp/article/details/52934933)
   
   - 当前项目[redis配置文件](./conf/redis.conf)
### redis操作命令  
   - #### 服务器端  
        - 服务器端的命令为  
        `redis-server`

        - 可以使⽤help查看帮助⽂档  
        `redis-server --help`

        - 推荐使⽤服务的⽅式管理redis服务  
        `启动: sudo service redis start`  
        `停⽌:sudo service redis stop`  
        `重启 sudo service redis restart`  
        - 其他  
        `ps -ef|grep redis 查看redis服务器进程`  
        `sudo kill -9 pid 杀死redis服务器`  
        `sudo redis-server /etc/redis/redis.conf 指定加载的配置文件`  

   - #### 客户端
        - 客户端的命令为redis-cli  
        
        - 可以使⽤help查看帮助⽂档  
        `redis-cli --help`

        - 连接redis  
        `redis-cli`

        - 切换数据库  
          数据库没有名称，默认有16个，通过0-15来标识，连接redis默认选择第一个数据库  
          `select n`
   
   
   
 