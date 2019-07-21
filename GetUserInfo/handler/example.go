package handler

import (
	"context"
	example "IHome/GetUserInfo/proto/example"
	"reflect"
	"github.com/astaxie/beego"
	"IHome/IhomeWeb/utils"
	"log"
	"IHome/IhomeWeb/model"
	"github.com/astaxie/beego/orm"
)

type Example struct{}


func (e *Example) GetUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("---------------- GET  /api/v1.0/user Getuserinfo() ------------------")
	//打印sessionid
	beego.Info(req.Sessionid,reflect.TypeOf(req.Sessionid))
	//错误码
	rsp.Errno  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Errno)


	//构建连接缓存的数据
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis connect err")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//拼接用户信息缓存字段
	sessioniduserid :=  req.Sessionid + "user_id"

	//获取到当前登陆用户的user_id
	value_id :=bm.Get(sessioniduserid)
	//打印
	beego.Info(value_id,reflect.TypeOf(value_id))

	//数据格式转换
	id :=  int(value_id.([]uint8)[0])
	beego.Info(id ,reflect.TypeOf(id))
	//创建user表
	user := models.User{Id:id}
	//创建数据库orm句柄
	o := orm.NewOrm()
	//查询表
	err =o.Read(&user)
	if err !=nil{
		rsp.Errno  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
		return  nil
	}
	//将查询到的数据依次赋值
	rsp.UserId= int64(user.Id)
	rsp.Name= user.Name
	rsp.Mobile = user.Mobile
	rsp.RealName = user.Real_name
	rsp.IdCard = user.Id_card
	rsp.AvatarUrl = user.Avatar_url
	return nil
}

