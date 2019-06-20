package handler

import (
	"context"

	example "IHome/PostUserAuth/proto/example"
	"fmt"
	"IHome/IhomeWeb/utils"
	"log"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"IHome/IhomeWeb/model"
	"time"
)

type Example struct{}

func (e *Example) PostUserAuth(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println(" 实名认证服务  PostUserAuth   /api/v1.0/user/auth  ")
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	//TODO 通过第三方接口进行实名认证调用，这里不做处理

	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis connect error")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	value := bm.Get(req.SessionId + "user_id")
	value_int, _ := redis.Int(value, nil)
	fmt.Println(value_int)
	//连接数据库
	o := orm.NewOrm()
	//更新数据
	var user = models.User{Id: value_int, Real_name: req.RealName, Id_card: req.IdCard}
	_, err = o.Update(&user, "real_name", "id_card")
	if err != nil {
		log.Println("mysql update error")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	o.Read(&user)

	//更新session信息时间
	bm.Put(req.SessionId+"user_id", user.Id, time.Second*600)
	bm.Put(req.SessionId+"mobile", user.Mobile, time.Second*600)
	bm.Put(req.SessionId+"name", user.Name, time.Second*600)
	return nil

}
