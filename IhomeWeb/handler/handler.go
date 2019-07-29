package handler

import (
	"context"
	"encoding/json"
	"net/http"
	GETAREA "IHome/GetArea/proto/example"
	GETIMAGECD "IHome/GetImageCd/proto/example"
	GETSMSCD "IHome/GetSmscd/proto/example"
	POSTRET "IHome/PostRet/proto/example"
	GETSESSION "IHome/GetSession/proto/example"
	POSTLOGIN "IHome/PostLogin/proto/example"
	DELETESESSION "IHome/DeleteSession/proto/example"
	GETUSERINFO "IHome/GetUserInfo/proto/example"
	POSTAVATAR "IHome/PostAvatar/proto/example"
	PUTUSERINFO "IHome/PutUserInfo/proto/example"
	POSTUSERAUTH "IHome/PostUserAuth/proto/example"
	GETUSERHOUSES "IHome/GetUserHouses/proto/example"
	POSTHOUSES "IHome/PostHouses/proto/example"
	POSTHOUSESIMAGE "IHome/PostHousesImage/proto/example"
	GETHOUSEINFO "IHome/GetHouseInfo/proto/example"
	GETINDEX "IHome/GetIndex/proto/example"
	GETHOUSES "IHome/GetHouses/proto/example"
	GETUSERORDER "IHome/GetUserOrder/proto/example"
	PUTORDERS "IHome/PutOrders/proto/example"
	POSTORDERS "IHome/PostOrders/proto/example"
	PUTCOMMENT "IHome/PutComment/proto/example"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	"github.com/astaxie/beego"
	"IHome/IhomeWeb/model"
	"image"
	"github.com/afocus/captcha"
	"image/png"
	"fmt"
	"regexp"
	"IHome/IhomeWeb/utils"
	"log"
	"io/ioutil"
)
//获取地区信息服务
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
//获取session信息服务
func GetSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取Session url：api/v1.0/session")
	cookie, err := r.Cookie("IHomelogin")
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
//获取首页轮播图的服务
func GetIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取首页轮播 url：api/v1.0/houses/index")
	server :=grpc.NewService()
	server.Init()

	exampleClient := GETINDEX.NewExampleService("go.micro.srv.GetIndex", server.Client())


	rsp, err := exampleClient.GetIndex(context.TODO(),&GETINDEX.Request{})
	if err != nil {
		beego.Info(err)
		http.Error(w, err.Error(), 502)
		return
	}
	data := []interface{}{}
	json.Unmarshal(rsp.Max,&data)

	//创建返回数据map
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":data,

	}
	w.Header().Set("Content-Type", "application/json")

	// 将返回数据map发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
}
//获取验证码图片服务
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
//获取短信验证码服务
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
//发送注册信息服务
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
	cookie, err := r.Cookie("IHomelogin")
	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{Name: "IHomelogin", Value: rsp.Sessionid, MaxAge: 600, Path: "/"}
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
//发送登陆信息服务
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
	cookie, err := r.Cookie("IHomelogin")
	log.Println(cookie)
	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{Name: "IHomelogin", Value: rsp.Sessionid, MaxAge: 400, Path: "/"}
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
//删除（退出）登陆信息服务
func DeleteSession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode the incoming request as json
	beego.Info("退出登陆 url：/api/v1.0/session Deletesession()")

	cli := grpc.NewService()
	cli.Init()
	// call the backend service
	exampleClient := DELETESESSION.NewExampleService("go.micro.srv.DeleteSession", cli.Client())
	//获取session
	userlogin, err := r.Cookie("IHomelogin")
	if err != nil || userlogin.Value == "" {
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
	if rsp.Errno == "0" {
		//将cookie中的sessionid设置为空
		_, err = r.Cookie("IHomelogin")
		if err == nil {
			cookie := http.Cookie{Name: "IHomelogin", Path: "/", MaxAge: -1}
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
//获取用户基本信息服务
func GetUserInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("获取用户信息 GetUserInfo /api/v1.0/user")
	cookie, err := r.Cookie("IHomelogin")
	if err != nil || cookie.Value == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
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
	exampleClient := GETUSERINFO.NewExampleService("go.micro.srv.GetUserInfo", cli.Client())

	//通过句柄调用服务端函数
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &GETUSERINFO.Request{
		Sessionid: cookie.Value,
	})

	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//准备返回数据
	data := make(map[string]interface{})
	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	// encode and write the response as json
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
//发送（上传）用户头像服务
func PostAvatar(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	beego.Info("上传用户touxiang PostAvatar /api/v1.0/user/avatar")
	//获取sessionid
	cookie, err := r.Cookie("IHomelogin")
	if err != nil || cookie.Value == "" {
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	file, header, err := r.FormFile("avatar")
	if err != nil {
		beego.Info("get file err:", err)
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	//文件校验
	filebuffer := make([]byte, header.Size)
	_, err = file.Read(filebuffer)
	if err != nil {
		beego.Info("get file err:", err)
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
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

	//通过protobuf生成文件创建连接服务端的客户端句柄
	exampleClient := POSTAVATAR.NewExampleService("go.micro.srv.PostAvatar", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.PostAvatar(context.TODO(), &POSTAVATAR.Request{
		Filename:  header.Filename,
		Filesize:  header.Size,
		Sessionid: cookie.Value,
		Avatar:    filebuffer,
	})

	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//返回数据
	data := make(map[string]interface{})
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	log.Println("data is ", data)
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
//更新用户名服务
func PutUserInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("更新用户名   PutUserInfo   /api/v1.0/user/name")
	//接受 前端发送过来数据的
	var request map[string]interface{}
	// 将前端 json 数据解析到 map当中
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//数据校验
	username := request["name"].(string)
	if username == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
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
	//获取sessionid
	cookie, err := r.Cookie("IHomelogin")
	if err != nil {
		log.Println("获取cookie失败")
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		// encode and write the response as json
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
	exampleClient := PUTUSERINFO.NewExampleService("go.micro.srv.PutUserInfo", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.PutUserInfo(context.TODO(), &PUTUSERINFO.Request{
		Sessionid: cookie.Value,
		Username:  request["name"].(string),
	})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//刷新cookie时间
	cookienew := http.Cookie{Name: "IHomelogin", Value: cookie.Value, Path: "/", MaxAge: 600}
	http.SetCookie(w, &cookienew)
	//返回数据
	data := make(map[string]interface{})
	data["name"] = rsp.Username
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  utils.RECODE_MOBILEERR,
		"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		"data":   data,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
//获取（检查）用户实名信息服务
func GetUserAuth(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("获取用户信息 GetUserInfo /api/v1.0/user")
	cookie, err := r.Cookie("IHomelogin")
	if err != nil || cookie.Value == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
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
	exampleClient := GETUSERINFO.NewExampleService("go.micro.srv.GetUserInfo", cli.Client())

	//通过句柄调用服务端函数
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &GETUSERINFO.Request{
		Sessionid: cookie.Value,
	})

	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//准备返回数据
	data := make(map[string]interface{})
	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	// encode and write the response as json
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
//发送用户实名认证信息服务
func PostUserAuth(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("更新实名认证检测  URL: /api/v1.0/user/auth PostUserAuth ")
	//接受 前端发送过来数据的
	var request map[string]interface{}
	// 将前端 json 数据解析到 map当中
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//数据校验
	if request["real_name"].(string) == "" || request["id_card"].(string) == "" {
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	//获取sessionid
	cookie, err := r.Cookie("IHomelogin")
	if err != nil || cookie.Value == "" {
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
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
	exampleClient := POSTUSERAUTH.NewExampleService("go.micro.srv.PostUserAuth", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.PostUserAuth(context.TODO(), &POSTUSERAUTH.Request{
		RealName:  request["real_name"].(string),
		IdCard:    request["id_card"].(string),
		SessionId: cookie.Value,
	})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//刷新cookie时间
	cookienew := http.Cookie{Name: "IHomelogin", Value: cookie.Value, Path: "/", MaxAge: 600}
	http.SetCookie(w, &cookienew)
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
//获取用户已发布房源信息服务
func GetUserHouses(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	beego.Info("获取当前用户所发布的房源 GetUserHouses /api/v1.0/user/houses")
	cli := grpc.NewService()
	cli.Init()
	// call the backend service
	exampleClient := GETUSERHOUSES.NewExampleService("go.micro.srv.GetUserHouses", cli.Client())
	userlogin, err := r.Cookie("IHomelogin")
	if err != nil || userlogin.Value != "" {
		//返回数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	rsp, err := exampleClient.GetUserHouses(context.TODO(), &GETUSERHOUSES.Request{
		Sessionid: userlogin.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//房屋切片信息
	house_list := []models.House{}
	json.Unmarshal(rsp.Mix, &house_list)
	//将房屋切片信息转换成map切片返回给前端
	var houses []interface{}
	for _, houseinfo := range house_list {
		houses = append(houses, houseinfo.To_house_info())
	}
	data_map := make(map[string]interface{})
	data_map["houses"] = houses
	//返回数据
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data_map,
	}

	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
//发送（发布）房源信息服务
func PostHouses(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// decode the incoming request as json
	beego.Info("PostHouses 发布房源信息 /api/v1.0/houses ")
	//获取前端post发送的请求信息
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("posthouses:",string(body))
	//获取cookie
	userlogin, err :=r.Cookie("ihomelogin")
	if err!=nil{
		resp := map[string]interface{}{
			"errno": utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}
	cli := grpc.NewService()
	cli.Init()
	// call the backend service
	exampleClient := POSTHOUSES.NewExampleService("go.micro.srv.PostHouses", cli.Client())
	rsp, err := exampleClient.PostHouses(context.TODO(), &POSTHOUSES.Request{
		Sessionid:userlogin.Value,
		Max:body,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		beego.Info(err)
		return
	}
	/*得到插入房源信息表的 id*/
	houseid_map:=make(map[string]interface{})
	houseid_map["house_id"]=int(rsp.HouseId)

	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   houseid_map,
	}

	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
//发送（上传）房屋图片服务
func PostHouseImage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	beego.Info("发送房屋图片PostHousesImage  /api/v1.0/houses/:id/images")
	//获取houseid
	houseid := params.ByName("id")
	fmt.Println("id")
	//获取sessionid
	userlogin, err := r.Cookie("ihomelogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}
	file, header, err := r.FormFile("house_image")
	if err != nil {
		beego.Info("Postupavatar   c.GetFile(avatar) err", err)

		resp := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}
	beego.Info(file, header)
	beego.Info("文件大小", header.Size)
	beego.Info("文件名", header.Filename)

	filebuffer := make([]byte, header.Size)
	_, err = file.Read(filebuffer)
	if err != nil {
		beego.Info("Postupavatar   file.Read(filebuffer) err", err)
		resp := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	cli := grpc.NewService()
	cli.Init()
	// call the backend service
	exampleClient := POSTHOUSESIMAGE.NewExampleService("go.micro.srv.PostHousesImage", cli.Client())
	rsp, err := exampleClient.PostHousesImage(context.TODO(), &POSTHOUSESIMAGE.Request{
		Sessionid: userlogin.Value,
		Id:        houseid,
		Image:     filebuffer,
		Filesize:  header.Size,
		Filename:  header.Filename,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//接收数据
	data := make(map[string]interface{})
	data["url"] = utils.AddDomain2Url(rsp.Url)
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}

	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
//获取房屋详细信息的服务
func GetHouseInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	beego.Info("获取房源详细信息 GetHouseInfo  api/v1.0/houses/:id ")

	//创建服务
	server :=grpc.NewService()
	server.Init()

	// call the backend service
	exampleClient := GETHOUSEINFO.NewExampleService("go.micro.srv.GetHouseInfo", server.Client())
	//获取房屋id
	id :=params.ByName("id")

	//获取sessionid
	userlogin,err:=r.Cookie("userlogin")
	if err != nil{
		resp := map[string]interface{}{
			"errno": utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}



	rsp, err := exampleClient.GetHouseInfo(context.TODO(), &GETHOUSEINFO.Request{
		//Sessionid
		Sessionid:userlogin.Value,
		//房屋id
		Id:id,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	//house := models.House{}
	house := make(map[string]interface{})

	json.Unmarshal(rsp.Housedata,&house)

	data_map :=make(map[string]interface{})
	//用户id
	data_map["user_id"] = int(rsp.Userid)
	//房屋详细信息
	data_map["house"] =  house


	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":data_map,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	return
}
//获取（搜索）房源服务
func GetHouses(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	//创建grpc
	server :=grpc.NewService()
	//初始化
	server.Init()

	exampleClient := GETHOUSES.NewExampleService("go.micro.srv.GetHouses", server.Client())

	//aid=5&sd=2017-11-12&ed=2017-11-30&sk=new&p=1
	aid := r.URL.Query()["aid"][0] //aid=5   地区编号
	sd := r.URL.Query()["sd"][0] //sd=2017-11-1   开始世界
	ed := r.URL.Query()["ed"][0] //ed=2017-11-3   结束世界
	sk := r.URL.Query()["sk"][0] //sk=new    第三栏条件
	p := r.URL.Query()["p"][0] //tp=1   页数

	rsp, err := exampleClient.GetHouses(context.TODO(), &GETHOUSES.Request{
		Aid:aid,
		Sd:sd,
		Ed:ed,
		Sk:sk,
		P:p,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}



	houses_l := []interface{}{}
	json.Unmarshal(rsp.Houses,&houses_l)

	data := map[string]interface{}{}
	data["current_page"] = rsp.CurrentPage
	data["houses"] = houses_l
	data["total_page"] = rsp.TotalPage

	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":data,
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}
//发布订单服务
func PostOrders(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	beego.Info("PostOrders  发布订单 /api/v1.0/orders")

	//将post代过来的数据转化以下
	body, _ := ioutil.ReadAll(r.Body)

	userlogin,err:=r.Cookie("userlogin")
	if err != nil||userlogin.Value==""{
		resp := map[string]interface{}{
			"errno": utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}


	service := grpc.NewService()
	service.Init()

	//调用服务
	exampleClient := POSTORDERS.NewExampleService("go.micro.srv.PostOrders", service.Client())
	rsp, err := exampleClient.PostOrders(context.TODO(), &POSTORDERS.Request{
		//sessionid
		Sessionid:userlogin.Value,
		//前端发送过来的数据
		Body:body,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	/*得到插入房源信息表的 id*/
	houseid_map :=make(map[string]interface{})
	houseid_map["order_id"] = int(rsp.OrderId)


	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":houseid_map,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}
//获取房东/租户订单信息服务
func GetUserOrder(w http.ResponseWriter, r *http.Request, params httprouter.Params) {


	beego.Info("/api/v1.0/user/orders   GetUserOrder 获取订单 ")
	server :=grpc.NewService()
	server.Init()
	// call the backend service
	exampleClient := GETUSERORDER.NewExampleService("go.micro.srv.GetUserOrder", server.Client())

	//获取cookie
	userlogin,err:=r.Cookie("userlogin")
	if err != nil{
		resp := map[string]interface{}{
			"errno": utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}
	//获取role
	role := r.URL.Query()["role"][0] //role


	rsp, err := exampleClient.GetUserOrder(context.TODO(), &GETUSERORDER.Request{
		//sessionid
		Sessionid:userlogin.Value,
		//角色
		Role:role,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	order_list := []interface{}{}
	json.Unmarshal(rsp.Orders,&order_list)

	data := map[string]interface{}{}
	data["orders"] = order_list



	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":data,
	}


	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}
//更新房东同意/拒绝订单
func PutOrders(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	// decode the incoming request as json
	//接收请求携带的数据    处理json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	//获取cookie   拿到cookie当中的数据
	userlogin,err:=r.Cookie("userlogin")
	if err != nil{
		resp := map[string]interface{}{
			"errno": utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 502)
			beego.Info(err)
			return
		}
		return
	}
	//创建grpc
	server:=grpc.NewService()
	//初始化
	server.Init()

	// call the backend service
	exampleClient := PUTORDERS.NewExampleService("go.micro.srv.PutOrders", server.Client())

	rsp, err := exampleClient.PutOrders(context.TODO(), &PUTORDERS.Request{
		//sessionid
		Sessionid:userlogin.Value,
		//具体操作
		Action:request["action"].(string),
		//订单id
		Orderid:params.ByName("id"),
	})
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 504)
		return
	}
}
//更新用户评价订单信息
func PutComment(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	beego.Info("PutComment  用户评价 /api/v1.0/orders/:id/comment")
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	service := grpc.NewService()
	service.Init()
	exampleClient := PUTCOMMENT.NewExampleService("go.micro.srv.PutComment", service.Client())

	//获取cookie
	userlogin,err:=r.Cookie("userlogin")
	if err != nil{
		resp := map[string]interface{}{
			"errno": utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	rsp, err := exampleClient.PutComment(context.TODO(), &PUTCOMMENT.Request{
		//sessionid
		Sessionid:userlogin.Value,
		//评价
		Comment:request["comment"].(string),
		//订单id
		OrderId:params.ByName("id"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}