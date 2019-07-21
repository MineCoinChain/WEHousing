package handler

import (
	"context"
	example "IHome/PutComment/proto/example"
	"reflect"
	"strconv"
	"github.com/astaxie/beego"
	"IHome/IhomeWeb/utils"
	"log"
	"IHome/IhomeWeb/model"
	"github.com/astaxie/beego/orm"
)

type Example struct{}


func (e *Example) PutComment(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("==============api/v1.0/orders  Postorders post succ!!=============")
	//创建返回空间
	rsp.Errno  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Errno)
	//1得到被评论的order_id
	//获得用户id
	//构建连接缓存的数据
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis connect err")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	sessioniduserid :=  req.Sessionid + "user_id"


	value_id :=bm.Get(sessioniduserid)
	beego.Info(value_id,reflect.TypeOf(value_id))
	user_id :=  int(value_id.([]uint8)[0])
	beego.Info(user_id ,reflect.TypeOf(user_id))

	//得到订单id
	order_id ,_ := strconv.Atoi(req.OrderId)
	//获得参数


	comment := req.Comment
	//检验评价信息是否合法 确保不为空
	if comment == "" {

		rsp.Errno  =  utils.RECODE_PARAMERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
		return nil
	}

	//2根据order_id找到所关联的房源信息
	//查询数据库，订单必须存在，订单状态必须为WAIT_COMMENT待评价状态
	order := models.OrderHouse{}
	o := orm.NewOrm()
	if err := o.QueryTable("order_house").Filter("id", order_id).Filter("status", models.ORDER_STATUS_WAIT_COMMENT).One(&order); err != nil {

		rsp.Errno  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)


		return nil
	}

	//关联查询order订单所关联的user信息
	if _, err := o.LoadRelated(&order, "User"); err != nil {

		rsp.Errno  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)

		return nil
	}

	//确保订单所关联的用户和该用户是同一个人
	if user_id != order.User.Id {


		rsp.Errno  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)


		return nil
	}

	//关联查询order订单所关联的House信息
	if _, err := o.LoadRelated(&order, "House"); err != nil {

		rsp.Errno  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)

		return nil
	}
	house := order.House
	//3将房源信息的评论字段追加评论信息
	//更新order的status为COMPLETE
	order.Status = models.ORDER_STATUS_COMPLETE
	order.Comment = comment

	//将房屋订单成交量+1
	house.Order_count++

	//将order和house更新数据库
	if _, err := o.Update(&order, "status", "comment"); err != nil {
		beego.Error("update order status, comment error, err = ", err)


		rsp.Errno  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)

		return nil
	}

	if _, err := o.Update(house, "order_count"); err != nil {
		beego.Error("update house order_count error, err = ", err)


		rsp.Errno  =  utils.RECODE_DATAERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)

		return nil
	}

	//将house_info_[house_id]的缓存key删除 （因为已经修改订单数量）
	house_info_key := "house_info_" + strconv.Itoa(house.Id)
	if err := bm.Delete(house_info_key); err != nil {
		beego.Error("delete ", house_info_key, "error , err = ", err)
	}

	return nil
}


