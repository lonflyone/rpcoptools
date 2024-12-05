package service

import (
	"context"
	"encoding/json"
	"net"
	"rpcoptool/app/task/repository/db/dao"
	"rpcoptool/app/task/repository/db/model"
	"rpcoptool/app/task/repository/mq"
	"rpcoptool/pkg/e"
	log "rpcoptool/pkg/logger"
	"rpcoptool/proto/pb"
	"sync"
)

var TaskService *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct{}

// 简单的单例模式
func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskService = &TaskSrv{}
	})
	return TaskService
}
func (T *TaskSrv) GetTasksList(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskListResponse) (err error) {
	log.LogrusObj.Infof("GetTasksList")
	resp.Code = e.SUCCESS
	if req.Limit == 0 {
		req.Limit = 10
	}
	r, count, err := dao.NewTaskDao(ctx).ListTask(int(req.Start), int(req.Limit))
	if err != nil {
		resp.Code = e.ERROR
		log.LogrusObj.Errorf("GetTasksList err:%v", err)
		return
	}
	// 返回proto里面定义的类型
	var taskRes []*pb.TaskModel
	for _, item := range r {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.Uip = getLocalIP()
	resp.TaskList = taskRes
	resp.Count = uint32(count)
	return
}
func (T *TaskSrv) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	log.LogrusObj.Infof("CreateTask")
	resp.Code = e.SUCCESS
	model := &model.Task{
		Id:        uint(req.Id),
		Title:     req.Title,
		Status:    int(req.Status),
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Url:       req.Url,
	}
	log.LogrusObj.Infof("Model %v", model)
	err = dao.NewTaskDao(ctx).CreateTask(model)
	if err != nil {
		resp.Code = e.ERROR
		log.LogrusObj.Errorf("NewTaskDao fail %v", err)
		return
	}
	resp.TaskDetail = BuildTask(model)
	return
}
func (T *TaskSrv) GetTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	log.LogrusObj.Infof("GetTask")
	resp.Code = e.SUCCESS
	r, err := dao.NewTaskDao(ctx).GetTaskById(req.Id)
	if err != nil {
		resp.Code = e.ERROR
		log.LogrusObj.Errorf("GetTask err:%v", err)
		return
	}
	task := BuildTask(r)
	//mq的rpc模式
	body, _ := json.Marshal(task)
	correlationID, callbackQueue, err := mq.SendMessage2MQ(body)
	if err != nil {
		resp.Code = e.ERROR
		log.LogrusObj.Errorf(" 发送mq信息失败 %v", err)
		return
	}
	err = mq.GetTaskOpMessage(ctx, callbackQueue, correlationID)
	if err != nil {
		task.Content = "rpc调用失败"
		resp.Code = e.ERROR
		log.LogrusObj.Errorf("GetTaskOpMessage fail %v", err)
		return
	}
	task.Content = "完成rpc调用"
	resp.TaskDetail = task
	return
}
func (T *TaskSrv) UpdateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	log.LogrusObj.Infof("UpdateTask")
	resp.Code = e.SUCCESS
	return
}
func (T *TaskSrv) DeleteTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	log.LogrusObj.Infof("DeleteTask")
	resp.Code = e.SUCCESS
	return
}
func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
func BuildTask(item *model.Task) *pb.TaskModel {
	taskModel := pb.TaskModel{
		Id:         uint64(item.Id),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
		Url:        item.Url,
	}
	return &taskModel
}
