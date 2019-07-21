package handler

import (
	"context"


	example "IHome/GetUserOrder/proto/example"
	"reflect"
	"github.com/astaxie/beego"
	"IHome/IhomeWeb/utils"
	"log"
	"github.com/astaxie/beego/orm"
	"IHome/IhomeWeb/model"
	"encoding/json"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetUserOrder(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("==============/api/v1.0/user/orders  GetOrders post succ!!=============")
	//创建返回空间
	rsp.Errno  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Errno)
	//根据session得到当前用户的user_id
	//构建连接缓存的数据
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis connect err")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//拼接key
	sessioniduserid :=  req.Sessionid + "user_id"

	//获取当前登陆的user_id
	value_id :=bm.Get(sessioniduserid)
	beego.Info(value_id,reflect.TypeOf(value_id))
	userid :=  int(value_id.([]uint8)[0])
	beego.Info(userid ,reflect.TypeOf(userid))

	//得到用户角色
	beego.Info(req.Role)

	//创建一个数据库对象
	o := orm.NewOrm()

	orders := []models.OrderHouse{}
	order_list := []interface{}{} //存放订单的切片

	if "landlord" == req.Role {
		//角色为房东
		//现找到自己目前已经发布了哪些房子
		landLordHouses := []models.House{}
		o.QueryTable("house").Filter("user__id", userid).All(&landLordHouses)

		housesIds := []int{}
		//遍历 自己所有的房子切片  得到所有房屋的id
		for _, house := range landLordHouses {
			housesIds = append(housesIds, house.Id)
		}
		//在从订单中找到房屋id为自己房源的id
		o.QueryTable("order_house").Filter("house__id__in", housesIds).OrderBy("ctime").All(&orders)
	} else {
		//角色为租客
		_,err:=o.QueryTable("order_house").Filter("user__id", userid).OrderBy("ctime").All(&orders)
		if err != nil {
			beego.Info(err)
		}

	}
	//循环将数据放到切片中
	for _, order := range orders {
		o.LoadRelated(&order, "User")
		o.LoadRelated(&order, "House")
		order_list = append(order_list,order.To_order_info())
	}

	rsp.Orders , _ = json.Marshal(order_list)

	return nil
}


