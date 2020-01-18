package main

import (
	"GoMicroServer/IHomeWeb/handler"
	"github.com/micro/go-micro/util/log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/web"

	// 数据库初始化
	_ "GoMicroServer/IHomeWeb/models"
)

func main() {
	// create new web service
	// 创建一个服务
	service := web.NewService(
		web.Name("go.micro.web.IhomeWeb"),
		web.Version("latest"),
		web.Address(":10086"),
	)

	// initialise service
	// 初始化服务
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// 使用路由器中间件映射路由
	rout := httprouter.New()
	rout.NotFound = http.FileServer(http.Dir("html"))

	//rout.GET("/api/v1.0/areas", )
	rout.GET("/api/v1.0/areas", handler.GetAreas)

	service.Handle("/", rout)

	// register html handler
	// service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	//service.HandleFunc("/Ihomeweb/call", handler.IhomewebCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
