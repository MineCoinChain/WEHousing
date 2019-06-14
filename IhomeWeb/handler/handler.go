package handler

import (
	"context"
	"encoding/json"
	"net/http"
	GETAREA "sss/GetArea/proto/example"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	"github.com/astaxie/beego"
	"sss/IhomeWeb/model"
)


func GetArea(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode the incoming request as json
	beego.Info("获取地区请求客户端 url：api/v1.0/areas")

	cli := grpc.NewService()
	cli.Init()
	// call the backend service
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea", cli.Client())
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//接收数据
	var areas []models.Area
	for _, value := range rsp.Data {
		temp := models.Area{Id: int(value.Aid), Name: value.Aname}
		areas=append(areas,temp)
	}
	response := map[string]interface{}{
		"errno":rsp.Errno,
		"errmsg":rsp.Errmsg,
		"data":areas,
	}

	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetSession(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {

	//准备返回给前端的map
	response := map[string]interface{}{
		"errno": "4101",
		"errmsg": "用户未登录",
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type","application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetIndex(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno": "0",
		"errmsg": "ok",
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type","application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}