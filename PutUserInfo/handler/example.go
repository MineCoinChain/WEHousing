package handler

import (
	"context"
	example "IHome/PutUserInfo/proto/example"
	"IHome/IhomeWeb/utils"
	"log"
	"github.com/garyburd/redigo/redis"
	"IHome/IhomeWeb/model"
	"github.com/astaxie/beego/orm"
	"time"
)

type Example struct{}

func (e *Example) PutUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	//初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	//连接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis连接错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	value := bm.Get(req.Sessionid + "user_id")
	userid, _ := redis.Int(value, nil)
	//更新数据库
	var user models.User
	user.Id = userid
	user.Name = req.Username
	o := orm.NewOrm()
	_, err = o.Update(&user, "Name")
	if err != nil {
		log.Println("数据库查询错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//修改redis中的信息
	err = o.Read(&user)
	if err != nil {
		log.Println("数据库查询错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	err = bm.Put(req.Sessionid+"user_id", user.Id, time.Second*600)
	if err != nil {
		log.Println("数据库查询错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	err = bm.Put(req.Sessionid+"name", user.Name, time.Second*600)
	if err != nil {
		log.Println("数据库查询错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	err = bm.Put(req.Sessionid+"mobile", user.Mobile, time.Second*600)
	if err != nil {
		log.Println("数据库查询错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	return nil

}
