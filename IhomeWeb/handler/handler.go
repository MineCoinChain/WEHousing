package handler

import (
	"context"
	"encoding/json"
	"net/http"
	GETAREA "sss/GetArea/proto/example"
	GETIMAGECD "sss/GetImageCd/proto/example"
	GETSMSCD "sss/GetSmscd/proto/example"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	"github.com/astaxie/beego"
	"sss/IhomeWeb/model"
	"image"
	"github.com/afocus/captcha"
	"image/png"
	"fmt"
	"regexp"
	"sss/IhomeWeb/utils"
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
		areas = append(areas, temp)
	}
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   areas,
	}

	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  "4101",
		"errmsg": "用户未登录",
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  "0",
		"errmsg": "ok",
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetImageCd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取图片验证码 url：/api/v1.0/imagecode/:uuid")
	uuid := ps.ByName("uuid")
	cli := grpc.NewService()
	cli.Init()
	// call the backend service
	exampleClient := GETIMAGECD.NewExampleService("go.micro.srv.GetImageCd", cli.Client())
	rsp, err := exampleClient.GetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid: uuid,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//判断是否返回图片
	if rsp.Errno != "0" {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"errno":  rsp.Errno,
			"errmsg": rsp.Errmsg,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	//拼接图片结构体发送给前端
	var img image.RGBA
	for _, value := range rsp.Pix {
		img.Pix = append(img.Pix, uint8(value))
	}
	img.Stride = int(rsp.Stride)
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)

	var image captcha.Image
	image.RGBA = &img
	fmt.Println(image)
	png.Encode(w, image)
}

func Getsmscd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	beego.Info("获取短信验证码 /api/v1.0/smscode/:mobile")
	//获取手机号
	mobile := ps.ByName("mobile")
	//验证手机号
	myreg := regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)
	bo := myreg.MatchString(mobile)
	if bo == false {
		response := map[string]interface{}{
			"errno":  utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	//获取图片验证码与uuid
	text := r.URL.Query()["text"][0]
	uuid := r.URL.Query()["id"][0]

	//创建客户端句柄，进行远程调用
	cli := grpc.NewService()
	cli.Init()
	exampleClient := GETSMSCD.NewExampleService("go.micro.srv.GetSmscd", cli.Client())
	rsp, err := exampleClient.GetSmscd(context.TODO(), &GETSMSCD.Request{
		Mobile:mobile,
		Text:text,
		Id:uuid,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//接收数据

	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
