package handler

import (
	"context"
	example "IHome/GetHouseInfo/proto/example"
	"reflect"
	"strconv"
	"time"
	"github.com/astaxie/beego"
	"IHome/IhomeWeb/utils"
	"encoding/json"
	"fmt"
	"IHome/IhomeWeb/model"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

func (e *Example) GetHouseInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("获取房源详细信息 GetHouseInfo  api/v1.0/houses/:id ")

	//创建返回空间
	rsp.Errno  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Errno)


	/*从session中获取我们的user_id的字段 得到当前用户id*/
	/*通过session 获取我们当前登陆用户的user_id*/
	//构建连接缓存的数据
	bm ,err :=utils.RedisOpen(utils.G_server_name,utils.G_redis_addr,
		utils.G_redis_port,utils.G_redis_dbnum)
	if err !=nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//拼接key
	sessioniduserid :=  req.Sessionid + "user_id"


	value_id :=bm.Get(sessioniduserid)
	beego.Info(value_id,reflect.TypeOf(value_id))
	//将[]uint8 强转为int
	id :=  int(value_id.([]uint8)[0])
	beego.Info(id ,reflect.TypeOf(id))

	/*从请求中的url获取房源id*/

	houseid,_ := strconv.Atoi(req.Id)

	/*从缓存数据库中获取到当前房屋的数据*/
	//缓存房屋信息的key是可变的
	house_info_key := fmt.Sprintf("house_info_%s", houseid)
	//从缓存中渠取道当前访问的信息
	house_info_value := bm.Get(house_info_key)
	if house_info_value!=nil{
		rsp.Userid= int64(id)
		rsp.Housedata= house_info_value.([]byte)
		return nil
	}


	/*查询当前数据库得到当前的house详细信息*/
	//创建数据对象
	house := models.House{Id:houseid}
	//创建数据库句柄
	o:= orm.NewOrm()
	//查询房屋非关系的基本信息
	o.Read(&house)
	/*关联查询 area user images fac等表*/
	o.LoadRelated(&house,"Area")
	o.LoadRelated(&house,"User")
	o.LoadRelated(&house,"Images")
	o.LoadRelated(&house,"Facilities")

	houseone :=house.To_one_house_desc()

	/*将查询到的结果存储到缓存当中*/
	housemix ,err :=json.Marshal(houseone)
	bm.Put(house_info_key,housemix,time.Second*3600)

	/*返回正确数据给前端*/

	rsp.Userid= int64(id)
	rsp.Housedata= housemix

	return nil
}


