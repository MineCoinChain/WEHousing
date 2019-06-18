package utils

import (
	"github.com/astaxie/beego/cache"
	"encoding/json"
	"log"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"crypto/md5"
	"encoding/hex"
	"github.com/weilaihui/fdfs_client"
)

/* 将url加上 http://IP:PROT/  前缀 */
//http:// + 127.0.0.1 + ：+ 8080 + 请求
//https://img.alicdn.com/tps/i4/TB1L7lExXzqK1RjSZFoSuvfcXXa.jpg_q90_.webp
func AddDomain2Url(url string) (domain_url string) {
	domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url

	return domain_url
}

//连接函数
func RedisOpen(server_name, redis_addr, redis_port, redis_dbnum string) (bm cache.Cache, err error) {
	redis_config_map := map[string]string{
		"key":   server_name,
		"conn":  redis_addr + ":" + redis_port,
		"dbnum": redis_dbnum,
	}
	redis_config, _ := json.Marshal(redis_config_map)
	bm, err = cache.NewCache("redis", string(redis_config))
	if err != nil {
		log.Println("连接redis错误", err)
		return nil, err
	}
	return bm, nil
}

//md5加密
func Getmd5string(s string) string {
	m := md5.New()
	return hex.EncodeToString(m.Sum([]byte(s)))
}

//按照按照二进制流上传文件
func Uploadbybuf(file []byte, extname string) (fileid string, err error) {
	//读取配置文件创建fdfs句柄
	fdfsclient, err := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")
	if err != nil {
		log.Fatal("fdfs_client.NewFdfsClient 创建失败", err)
		return "", err
	}

	//上传文件
	rsp, err := fdfsclient.UploadByBuffer(file, extname)
	if err != nil {
		log.Println("fdfsclient.UploadByBuffer 上传失败",err)
		return "",err
	}
	//返回filedid
	return rsp.RemoteFileId,nil
}
