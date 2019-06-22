package main

import (
        "github.com/micro/go-log"
        "github.com/micro/go-web"
        "github.com/julienschmidt/httprouter"
        _ "IHome/IhomeWeb/model"

        "net/http"
        "IHome/IhomeWeb/handler"
)

func main() {
	    // 构造web服务
        service := web.NewService(
                web.Name("go.micro.web.IhomeWeb"),
                web.Version("latest"),
                web.Address(":22333"),
        )
	    // 服务初始化
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }
        //构建路由
		rou:=httprouter.New()

		//将路由注册到服务
		rou.NotFound = http.FileServer(http.Dir("html"))
        rou.GET("/api/v1.0/areas",handler.GetArea)
        //欺骗浏览器  session index
        rou.GET("/api/v1.0/session", handler.GetSession)
        //session
        rou.GET("/api/v1.0/house/index", handler.GetIndex)
        //获取图片验证码
        rou.GET("/api/v1.0/imagecode/:uuid",handler.GetImageCd)
        //获取短信验证码
        rou.GET("/api/v1.0/smscode/:mobile",handler.Getsmscd)
        //用户注册
        rou.POST("/api/v1.0/users",handler.PostRet)
        //用户登陆
        rou.POST("/api/v1.0/sessions",handler.PostLogin)
        //退出登陆
        rou.DELETE("/api/v1.0/session", handler.DeleteSession)
        //获取用户详细信息
        rou.GET("/api/v1.0/user",handler.GetUserInfo)
        //用户上传图片
        rou.POST("/api/v1.0/user/avatar",handler.PostAvatar)
        //请求更新用户名
        rou.PUT("/api/v1.0/user/name",handler.PutUserInfo)
        //身份认证检查 同  获取用户信息   所调用的服务是 GetUserInfo
        rou.GET("/api/v1.0/user/auth", handler.GetUserAuth)
        //实名认证服务
        rou.POST("/api/v1.0/user/auth",handler.PostUserAuth)
        //获取用户已发布房源信息服务
        rou.GET("/api/v1.0/user/houses",handler.GetUserHouses)
        //发送（发布）房源信息服务
        rou.POST("/api/v1.0/houses",handler.PostHouses)
        //发送（上传）房屋图片服务
        rou.POST("/api/v1.0/houses/:id/images",handler.PostHouseImage)
		service.Handle("/",rou)
		// 服务运行
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
