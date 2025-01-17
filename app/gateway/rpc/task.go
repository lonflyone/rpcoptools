package rpc

import (
	"context"
	"rpcoptool/pkg/e"
	"rpcoptool/proto/pb"
)

func TaskList(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskListResponse, err error) {
	TaskService := TaskServices.Get().(pb.TaskService)
	defer TaskServices.Put(TaskService)

	r, err := TaskService.GetTasksList(ctx, req)
	if err != nil {
		return
	}
	if r.Code != e.SUCCESS {
		return
	}

	return r, nil
}
func TaskGet(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	TaskService := TaskServices.Get().(pb.TaskService)
	defer TaskServices.Put(TaskService)
	r, err := TaskService.GetTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != e.SUCCESS {
		return
	}

	return r, nil
}
func TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	TaskService := TaskServices.Get().(pb.TaskService)
	defer TaskServices.Put(TaskService)
	r, err := TaskService.CreateTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != e.SUCCESS {
		return
	}

	return r, nil
}

func TaskUpdate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	TaskService := TaskServices.Get().(pb.TaskService)
	defer TaskServices.Put(TaskService)
	r, err := TaskService.UpdateTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != e.SUCCESS {
		return
	}

	return r, nil
}

func TaskDelete(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	TaskService := TaskServices.Get().(pb.TaskService)
	defer TaskServices.Put(TaskService)
	r, err := TaskService.DeleteTask(ctx, req)
	if err != nil {
		return
	}
	if r.Code != e.SUCCESS {
		return
	}

	return r, nil
}
