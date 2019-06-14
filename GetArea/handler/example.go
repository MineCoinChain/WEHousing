package handler

import (
	"context"
	example "sss/GetArea/proto/example"
	"fmt"
	"sss/IhomeWeb/utils"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/model"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取地域信息服务   GetArea  /api/v1.0/areas")
	//1.初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	//查询数据库
	o := orm.NewOrm()
	//接受数据
	var areas []models.Area
	//设置查询条件
	qs := o.QueryTable("area")
	//查询全部
	num, err := qs.All(&areas)
	if err != nil {
		fmt.Println("查询数据库错误")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if num == 0 {
		fmt.Println("无数据", err)
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//4.将查询到的数据传入protobuf中
	for _, value := range areas {
		temp := example.Response_Address{}
		temp.Aid = int32(value.Id)
		temp.Aname = value.Name
		rsp.Data = append(rsp.Data, &temp)
	}
	return nil
}
