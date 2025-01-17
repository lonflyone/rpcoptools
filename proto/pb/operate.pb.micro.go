// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: operate.proto

package pb

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for OperatorService service

func NewOperatorServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OperatorService service

type OperatorService interface {
	CreateTaskOperator(ctx context.Context, in *TaskReq, opts ...client.CallOption) (*TaskOperatorInfo, error)
}

type operatorService struct {
	c    client.Client
	name string
}

func NewOperatorService(name string, c client.Client) OperatorService {
	return &operatorService{
		c:    c,
		name: name,
	}
}

func (c *operatorService) CreateTaskOperator(ctx context.Context, in *TaskReq, opts ...client.CallOption) (*TaskOperatorInfo, error) {
	req := c.c.NewRequest(c.name, "OperatorService.CreateTaskOperator", in)
	out := new(TaskOperatorInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OperatorService service

type OperatorServiceHandler interface {
	CreateTaskOperator(context.Context, *TaskReq, *TaskOperatorInfo) error
}

func RegisterOperatorServiceHandler(s server.Server, hdlr OperatorServiceHandler, opts ...server.HandlerOption) error {
	type operatorService interface {
		CreateTaskOperator(ctx context.Context, in *TaskReq, out *TaskOperatorInfo) error
	}
	type OperatorService struct {
		operatorService
	}
	h := &operatorServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OperatorService{h}, opts...))
}

type operatorServiceHandler struct {
	OperatorServiceHandler
}

func (h *operatorServiceHandler) CreateTaskOperator(ctx context.Context, in *TaskReq, out *TaskOperatorInfo) error {
	return h.OperatorServiceHandler.CreateTaskOperator(ctx, in, out)
}
