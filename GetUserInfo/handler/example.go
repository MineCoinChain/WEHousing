package handler

import (
	"context"

	example "IHome/GetUserInfo/proto/example"
	"IHome/IhomeWeb/utils"
	"log"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"IHome/IhomeWeb/model"
	"strconv"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	//连接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis connect failed")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	value := bm.Get(req.Sessionid + "user_id")
	value_int, _ := redis.Int(value, nil)
	//查询数据库
	var o = orm.NewOrm()
	var user models.User
	user.Id = value_int
	err = o.Read(&user)
	if err != nil {
		log.Println("mysql err:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	rsp.UserId = strconv.Itoa(user.Id)
	rsp.Name = user.Name
	rsp.Mobile = user.Mobile
	rsp.AvatarUrl = user.Avatar_url
	rsp.IdCard = user.Id_card
	rsp.RealName = user.Real_name
	return nil

}
