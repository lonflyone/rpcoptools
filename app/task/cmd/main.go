package main

import (
	"fmt"
	"rpcoptool/app/task/repository/db/dao"
	"rpcoptool/app/task/repository/mq"
	"rpcoptool/app/task/service"
	log "rpcoptool/pkg/logger"
	"rpcoptool/proto/pb"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	dao.InitDB()
	log.InitLog()
	mq.InitRabbitMQ()
	EtcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", "localhost", "2379")),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcTaskService"), // 微服务名字
		// micro.Address(":8084"),通过名字找到，倒也不用实际的端口
		micro.Registry(EtcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterTaskServiceHandler(microService.Server(), service.GetTaskSrv())
	// 启动微服务
	_ = microService.Run()
}
