package main

import (
        "github.com/micro/go-log"
        "github.com/micro/go-web"
        "github.com/julienschmidt/httprouter"
        _ "sss/IhomeWeb/model"

        "net/http"
        "sss/IhomeWeb/handler"
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

		service.Handle("/",rou)
		// 服务运行
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
