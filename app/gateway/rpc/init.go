package rpc

import (
	"fmt"
	"rpcoptool/app/gateway/wrappers"
	"rpcoptool/proto/pb"
	"sync"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

// pb.TaskService
var (
	TaskServices *sync.Pool
	EtcdReg      registry.Registry
)

func InitRPC() {
	EtcdReg = etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", "localhost", "2379")),
	)
	// task 连接池
	TaskServices = &sync.Pool{
		New: func() any {
			// task
			taskMicroService := micro.NewService(
				micro.Registry(EtcdReg),
				micro.Name("taskService.client"),
				micro.WrapClient(wrappers.NewTaskWrapper),
				micro.Selector(selector.NewSelector(
					selector.SetStrategy(selector.RoundRobin),
					selector.Registry(EtcdReg),
				)),
			)
			return pb.NewTaskService("rpcTaskService", taskMicroService.Client())
		},
	}

}
