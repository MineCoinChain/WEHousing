package main

import (
        "github.com/micro/go-log"
        "github.com/micro/go-web"
        "github.com/julienschmidt/httprouter"
        _ "sss/IhomeWeb/model"

        "net/http"
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
		//rou.GET("/example/call", handler.ExampleCall)
		//将路由注册到服务
		rou.NotFound = http.FileServer(http.Dir("html"))
		service.Handle("/",rou)
		// 服务运行
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
