package handler

import (
	"context"
	example "IHome/PostOrders/proto/example"
	"reflect"
	"time"
	"strconv"
	"github.com/astaxie/beego"
	"IHome/IhomeWeb/utils"
	"log"
	"encoding/json"
	"IHome/IhomeWeb/model"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

func (e *Example) PostOrders(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("api/v1.0/orders  Postorders  发布订单=============")
	//创建返回空间
	rsp.Errno  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Errno)

	//1根据session得到当前用户的user_id
	//构建连接缓存的数据
	//连接redis
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis connect err:", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//拼接key
	sessioniduserid :=  req.Sessionid + "user_id"


	value_id :=bm.Get(sessioniduserid)
	beego.Info(value_id,reflect.TypeOf(value_id))
	userid :=  int(value_id.([]uint8)[0])
	beego.Info(userid ,reflect.TypeOf(userid))

	//2得到用户请求的json数据并效验合法性
	//获取用户请求Response数据的name
	var RequestMap = make(map[string]interface{})

	err  =json.Unmarshal(req.Body, &RequestMap)
	if err != nil {

		rsp.Errno  =  utils.RECODE_REQERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
		return nil
	}
	beego.Info(RequestMap)

	//效验合法性
	//用户参数做合法判断
	if RequestMap["house_id:"]== "" || RequestMap["start_date"] == "" || RequestMap["end_date"] == "" {
		rsp.Errno  =  utils.RECODE_REQERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
		return nil
	}


	//3确定end_date在start_data之后
	//格式化日期时间
	start_date_time, _ := time.Parse("2006-01-02 15:04:05",RequestMap["start_date"].(string)+" 00:00:00")
	end_date_time, _ := time.Parse("2006-01-02 15:04:05", RequestMap["end_date"].(string)+" 00:00:00")

	//4得到一共入住的天数

	beego.Info(start_date_time,end_date_time)
	days := end_date_time.Sub(start_date_time).Hours()/24 + 1
	beego.Info( days)

	//5根据order_id得到关联的房源信息

	house_id, _ := strconv.Atoi(RequestMap["house_id"].(string))
	//房屋对象
	house := models.House{Id: house_id}
	o := orm.NewOrm()
	if err := o.Read(&house); err != nil {
		rsp.Errno  =  utils.RECODE_NODATA
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
		return nil
	}
	o.LoadRelated(&house, "User")

	//6确保当前的uers_id不是房源信息所关联的user_id
	if userid == house.User.Id {


		rsp.Errno  =  utils.RECODE_ROLEERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)

		return nil
	}
	//7确保用户选择的房屋未被预定,日期没有冲突
	if end_date_time.Before(start_date_time) {

		rsp.Errno  =  utils.RECODE_ROLEERR
		rsp.Errmsg  = "结束时间在开始时间之前"
		return nil
	}
	//7.1添加征信步骤
	//这一步通过第三方平台添加

	//8封装order订单
	amount := days * float64(house.Price)
	order := models.OrderHouse{}
	order.House = &house
	user := models.User{Id: userid}
	order.User = &user
	order.Begin_date = start_date_time
	order.End_date = end_date_time
	order.Days = int(days)
	order.House_price = house.Price
	order.Amount = int(amount)
	order.Status = models.ORDER_STATUS_WAIT_ACCEPT
	//征信
	order.Credit = false

	beego.Info(order)
	//9将订单信息入库表中
	if _, err := o.Insert(&order); err != nil {
		rsp.Errno  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
		return nil
	}
	//10返回order_id
	bm.Put(sessioniduserid, string(userid) ,time.Second*7200)
	rsp.OrderId = int64(order.Id)
	return nil

}


