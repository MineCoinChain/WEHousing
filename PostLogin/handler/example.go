package handler

import (
	"context"
	example "IHome/PostLogin/proto/example"
	"IHome/IhomeWeb/utils"
	"github.com/astaxie/beego/orm"
	"fmt"
	"IHome/IhomeWeb/model"
	"log"
	"strconv"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostLogin(ctx context.Context, req *example.Request, rsp *example.Response) error {
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	var user models.User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	err := qs.Filter("mobile", req.Mobile).One(&user)
	if err != nil {
		fmt.Println("用户名查询失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if utils.Getmd5string(req.Password) != user.Password_hash {
		log.Println("密码错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	sessionid := utils.Getmd5string(req.Mobile + req.Password + strconv.Itoa(int(time.Now().UnixNano())))
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("连接redis失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	bm.Put(sessionid+"name", user.Name, time.Second*600)
	bm.Put(sessionid+"mobile", user.Mobile, time.Second*600)
	bm.Put(sessionid+"user_id", user.Id, time.Second*600)
	rsp.Sessionid = sessionid
	return nil
}
