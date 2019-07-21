package handler

import (
	"context"


	example "IHome/GetIndex/proto/example"
	"time"
	"IHome/IhomeWeb/utils"
	"IHome/IhomeWeb/model"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"encoding/json"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetIndex(ctx context.Context, req *example.Request, rsp *example.Response) error {
	//创建返回空间
	rsp.Errno  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Errno)


	data := []interface{}{}
	//1 从缓存服务器中请求 "home_page_data" 字段,如果有值就直接返回
	//先从缓存中获取房屋数据,将缓存数据返回前端即可
	bm ,err :=utils.RedisOpen(utils.G_server_name,utils.G_redis_addr,
		utils.G_redis_port,utils.G_redis_dbnum)
	if err !=nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	house_page_key := "home_page_data"
	house_page_value := bm.Get(house_page_key)
	if house_page_value != nil {
		beego.Debug("======= get house page info  from CACHE!!! ========")
		//直接将二进制发送给客户端
		rsp.Max = house_page_value.([]byte)
		return nil
	}

	houses := []models.House{}

	//2 如果缓存没有,需要从数据库中查询到房屋列表
	o := orm.NewOrm()

	if _, err := o.QueryTable("house").Limit(models.HOME_PAGE_MAX_HOUSES).All(&houses); err == nil {
		for _, house := range houses {
			o.LoadRelated(&house, "Area")
			o.LoadRelated(&house, "User")
			o.LoadRelated(&house, "Images")
			o.LoadRelated(&house, "Facilities")
			data = append(data, house.To_house_info())
		}

	}
	beego.Info(data,houses)
	//将data存入缓存数据
	house_page_value, _ = json.Marshal(data)
	bm.Put(house_page_key, house_page_value, 3600*time.Second)

	rsp.Max= house_page_value.([]byte)
	return nil

}

