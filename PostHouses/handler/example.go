package handler

import (
	"context"

	example "IHome/PostHouses/proto/example"
	"github.com/astaxie/beego"
	"IHome/IhomeWeb/utils"
	"encoding/json"
	"IHome/IhomeWeb/model"
	"strconv"
	"log"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

func (e *Example) PostHouses(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("PostHouses 发布房源信息 /api/v1.0/houses ")
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	var Requestmap = make(map[string]interface{})
	json.Unmarshal(req.Max, &Requestmap)
	//构造house结构体
	house := models.House{}
	house.Title = Requestmap["title"].(string)
	price, _ := strconv.Atoi(Requestmap["price"].(string))
	house.Price = price * 100
	house.Address = Requestmap["address"].(string)
	house.Room_count, _ = strconv.Atoi(Requestmap["room_count"].(string))
	house.Acreage, _ = strconv.Atoi(Requestmap["acreage"].(string))
	house.Unit = Requestmap["unit"].(string)
	house.Capacity, _ = strconv.Atoi(Requestmap["capacity"].(string))
	house.Beds, _ = Requestmap["beds"].(string)
	depsoit, _ := strconv.Atoi(Requestmap["deposit"].(string))
	house.Deposit = depsoit * 100
	house.Min_days, _ = strconv.Atoi(Requestmap["min_days"].(string))
	house.Max_days, _ = strconv.Atoi(Requestmap["max_days"].(string))
	facility := []*models.Facility{}
	for _, f_id := range Requestmap["facility"].([]interface{}) {
		fid, _ := strconv.Atoi(f_id.(string))
		fac := &models.Facility{Id: fid}
		facility = append(facility, fac)
	}
	area_id, _ := strconv.Atoi(Requestmap["area_id"].(string))
	area := models.Area{Id: area_id}
	house.Area = &area
	//连接redis查询userid
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis 连接失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	value_id := bm.Get(req.Sessionid + "user_id")
	user_id, _ := redis.Int(value_id, nil)
	var o = orm.NewOrm()
	user := models.User{Id: user_id}
	o.Read(&user)
	house.User = &user
	houseid, err := o.Insert(&house)
	if err != nil {

		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	m2m := o.QueryM2M(&house, "Facilities")
	num, err := m2m.Add(facility)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if num == 0 {
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	rsp.HouseId = int64(houseid)
	return nil
}
