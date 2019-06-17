package handler

import (
	"context"
	"encoding/json"
	"net/http"
	GETAREA "sss/GetArea/proto/example"
	GETIMAGECD "sss/GetImageCd/proto/example"
	GETSMSCD "sss/GetSmscd/proto/example"
	POSTRET "sss/PostRet/proto/example"
	GETSESSION "sss/GetSession/proto/example"
	POSTLOGIN "sss/PostLogin/proto/example"
	DELETESESSION "sss/DeleteSession/proto/example"
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
	"log"
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
	beego.Info("获取Session url：api/v1.0/session")
	cookie, err := r.Cookie("ihomelogin")
	if err != nil {
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  "4101",
			"errmsg": "用户未登录",
		}
		// encode and write the response as json
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//准备返回给前端的map
	cli := grpc.NewService()
	cli.Init()

	exampleClient := GETSESSION.NewExampleService("go.micro.srv.GetSession", cli.Client())
	rsp, err := exampleClient.GetSession(context.TODO(), &GETSESSION.Request{
		Sessionid: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := make(map[string]string)
	data["name"] = rsp.Data
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
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
		Mobile: mobile,
		Text:   text,
		Id:     uuid,
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

func PostRet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" 注册服务  PostRet  /api/v1.0/users")
	//接受 前端发送过来数据的
	var request map[string]interface{}
	// 将前端 json 数据解析到 map当中
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if request["mobile"].(string) == "" || request["password"].(string) == "" || request["sms_code"].(string) == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	//创建 grpc 客户端
	cli := grpc.NewService()
	//客户端初始化
	cli.Init()

	//通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := POSTRET.NewExampleService("go.micro.srv.PostRet", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.PostRet(context.TODO(), &POSTRET.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
		SmsCode:  request["sms_code"].(string),
	})

	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//设置cookie
	cookie, err := r.Cookie("ihomelogin")
	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{Name: "ihomelogin", Value: rsp.Sessionid, MaxAge: 600, Path: "/"}
		http.SetCookie(w, &cookie)
	}
	//将数据返回前端
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func PostLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("登陆 api/v1.0/sessions")
	//接受 前端发送过来数据的
	var request map[string]interface{}
	// 将前端 json 数据解析到 map当中
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//校验数据
	if request["mobile"].(string) == "" || request["password"].(string) == "" {
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// encode and write the response as json
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		return
	}
	//创建 grpc 客户端
	cli := grpc.NewService()
	//客户端初始化
	cli.Init()

	//通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := POSTLOGIN.NewExampleService("go.micro.srv.PostLogin", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.PostLogin(context.TODO(), &POSTLOGIN.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
	})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	cookie, err := r.Cookie("ihomelogin")
	log.Println(cookie)
	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{Name: "ihomelogin", Value: rsp.Sessionid, MaxAge: 400, Path: "/"}
		http.SetCookie(w, &cookie)
	}
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func DeleteSession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode the incoming request as json
	beego.Info("退出登陆 url：/api/v1.0/session Deletesession()")

	cli := grpc.NewService()
	cli.Init()
	// call the backend service
	exampleClient := DELETESESSION.NewExampleService("go.micro.srv.DeleteSession", cli.Client())
	//获取session
	userlogin, err := r.Cookie("ihomelogin")
	if err != nil || userlogin.Value==""{
		log.Println("user not login")
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	rsp, err := exampleClient.DeleteSession(context.TODO(), &DELETESESSION.Request{
		Sessionid: userlogin.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if rsp.Errno=="0"{
		//将cookie中的sessionid设置为空
		_, err = r.Cookie("ihomelogin")
		if err == nil {
			cookie := http.Cookie{Name: "ihomelogin", Path: "/", MaxAge: -1}
			http.SetCookie(w, &cookie)
		}
	}
	//返回数据
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	//设置格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
