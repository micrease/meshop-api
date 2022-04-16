package main

import (
	"flag"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/web"
	nacos "github.com/micrease/micrease-core/registry"
	"meshop-api/app/service/remote_service"
	sysConfig "meshop-api/config"
	"meshop-api/router"
)

func InitServer() {
	//解析命令运行参数
	flag.Parse()
	//从config.json加载配置信息
	//加载环境配置
	sysConfig.InitSysConfig()
}

func main() {
	InitServer()
	//获取业务配置
	conf := sysConfig.Get()
	//定义gin router
	ginRouter := router.InitGinRouter()
	//注册中心
	nacosRegistry := nacos.NewRegistry(registry.Addrs(conf.Nacos.Addrs))
	//接收http请求,因此创建一个webService类型的服务,用gin框架做为路由
	httpServer := web.NewService(
		web.Address(":"+conf.Service.Port),
		web.Handler(ginRouter),
		web.Name(conf.Service.ServiceName),
		web.Version(conf.Service.Version),
		web.Registry(nacosRegistry),
	)
	if err := httpServer.Init(); err != nil {
		log.Fatal(err)
	}
	//创建RPC调用client
	remote_service.Register()
	// Run service
	if err := httpServer.Run(); err != nil {
		log.Fatal(err)
	}
}
