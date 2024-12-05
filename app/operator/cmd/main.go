package main

import (
	"context"
	"fmt"
	"rpcoptool/app/operator/repository/mq"
	"rpcoptool/app/operator/service"
	log "rpcoptool/pkg/logger"
	"rpcoptool/proto/pb"

	"rpcoptool/app/operator/repository/mq/operator"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	log.InitLog()
	mq.InitRabbitMQ()
	loadingScript()
	EtcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", "localhost", "2379")),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcOperatorService"), // 微服务名字
		micro.Registry(EtcdReg),          // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterOperatorServiceHandler(microService.Server(), service.GetOpSrv())
	// 启动微服务
	_ = microService.Run()
}
func loadingScript() {
	ctx := context.Background()
	//因为这个进程会被block住,所以起一个携程
	synctask := operator.SyncTask{}
	go synctask.RunTaskCreate(ctx)
}
