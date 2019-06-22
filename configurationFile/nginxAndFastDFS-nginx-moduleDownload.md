# nginx和fastDFS-nginx-module
- [mod_fastdfs.conf文件](./conf/mod_fastdfs.conf)
- [nginx.conf文件](./conf/nginx.conf)
- ### nginx和fastDFS-nginx-module安装
    1. 解压缩 nginx-1.8.1.tar.gz
    2. 解压缩 fastdfs-nginx-module-master.zip
    3. 进入nginx-1.8.1目录中
    4. 执行  
       `sudo ./configure --prefix=/usr/local/nginx/ --add-module=fastdfs-nginx-module-master解压后的目录的绝对路径/src`
       如:  
       `sudo ./configure --prefix=/usr/local/nginx/ --add-module=/home/hyw/下载/fastdfs-nginx-module-master/src`
    5. `sudo make`
    6. `sudo make install`
    7. `sudo cp fastdfs-nginx-module-master解压后的目录中src下的mod_fastdfs.conf  /etc/fdfs/mod_fastdfs.conf`
    8. 修改配置文件：[mod_fastdfs.conf文件](./conf/mod_fastdfs.conf)  
       `sudo vim /etc/fdfs/mod_fastdfs.conf`  
       修改内容：  
       `connect_timeout=10`  
       `tracker_server=自己ubuntu虚拟机的ip地址:22122`  
       `url_have_group_name=true`  
       `store_path0=/home/python/fastdfs/storage`
    7. `sudo cp 解压缩的fastdfs-master/conf/http.conf  /etc/fdfs/http.conf`
    8. `sudo cp 解压缩的fastdfs-master/conf/mime.types /etc/fdfs/mime.types`
    9. `sudo vim /usr/local/nginx/conf/nginx.conf`  
        在http部分中添加配置信息如下：[nginx.conf文件](./conf/nginx.conf)
        ```
        server {
            listen       8888;
            server_name  localhost;
            location ~/group[0-9]/ {
                ngx_fastdfs_module;
            }
            error_page   500 502 503 504  /50x.html;
            location = /50x.html {
            root   html;
            }
        }
        ```
    10. 启动nginx  
        `sudo /usr/local/nginx/sbin/nginx`
- ### 可能出现的错误
    - 安装nginx时遇到的包问题
       https://blog.csdn.net/z920954494/article/details/52132125
    - Nginx编译时报错
       https://blog.csdn.net/u010889616/article/details/82867091
    
