package handler

import (
	"context"
	example "IHome/PostAvatar/proto/example"
	"IHome/IhomeWeb/utils"
	"path"
	"log"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"IHome/IhomeWeb/model"
)

type Example struct{}

func (e *Example) PostAvatar(ctx context.Context, req *example.Request, rsp *example.Response) error {
	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.AddDomain2Url(rsp.Errno)
	//文件校验
	if len(req.Avatar) != int(req.Filesize) {
		log.Println("文件大小出错")
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg = utils.AddDomain2Url(rsp.Errno)
		return nil
	}
	//上传文件
	ext := path.Ext(req.Filename)
	fileid, err := utils.Uploadbybuf(req.Avatar, ext[1:])
	if err != nil {
		log.Println("文件上传出错:", err)
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg = utils.AddDomain2Url(rsp.Errno)
		return nil
	}
	//连接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis connect err:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//获取userid
	value := bm.Get(req.Sessionid + "user_id")
	userid, _ := redis.Int(value, nil)
	//更新mysql数据库
	o := orm.NewOrm()
	var user models.User
	user.Id = userid
	user.Avatar_url = fileid
	_, err = o.Update(&user, "avatar_url")
	if err != nil {
		log.Println("Mysql Update err", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	rsp.AvatarUrl = fileid
	return nil
}
