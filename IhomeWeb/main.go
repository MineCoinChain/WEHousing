package main

import (
        "github.com/micro/go-log"
	"net/http"

        "github.com/micro/go-web"
        "sss/IhomeWeb/handler"
        "github.com/julienschmidt/httprouter"
)

func main() {
	// 构造web服务
        service := web.NewService(
                web.Name("go.micro.web.IhomeWeb"),
                web.Version("latest"),
        )

	// 服务初始化
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

        //构建路由
		rou:=httprouter.New()
		rou.GET("/example/call", handler.ExampleCall)
		//将路由注册到服务
		service.Handle("/",rou)
		// 服务运行
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
