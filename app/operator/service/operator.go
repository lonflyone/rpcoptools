package service

import (
	"context"
	"rpcoptool/proto/pb"
	"sync"
)

var OperatorService *OpSrv
var OperatorSrvOnce sync.Once

type OpSrv struct{}

// 简单的单例模式
func GetOpSrv() *OpSrv {
	OperatorSrvOnce.Do(func() {
		OperatorService = &OpSrv{}
	})
	return OperatorService
}
func (o *OpSrv) CreateTaskOperator(ctx context.Context, req *pb.TaskReq, resp *pb.TaskOperatorInfo) (err error) {

	return
}
