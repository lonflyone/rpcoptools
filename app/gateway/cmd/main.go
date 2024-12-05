package main

import (
	"rpcoptool/app/gateway/rpc"
	"time"

	"rpcoptool/app/gateway/router"

	log "rpcoptool/pkg/logger"

	"go-micro.dev/v4/web"
)

func main() {
	//连接池方式
	rpc.InitRPC()
	log.InitLog()

	// 创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		// 将服务调用实例使用gin处理
		web.Handler(router.NewRouter()),
		web.Registry(rpc.EtcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	// 接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
