package http

import (
	"net/http"
	log "rpcoptool/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"

	"rpcoptool/app/gateway/rpc"
	"rpcoptool/pkg/ctl"
	"rpcoptool/proto/pb"
)

func ListTaskHandler(ctx *gin.Context) {
	log.LogrusObj.Infof("ListTaskHandler")
	var taskReq pb.TaskRequest
	if err := ctx.Bind(&taskReq); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数失败"))
		return
	}
	// 调用服务端的函数
	taskResp, err := rpc.TaskList(ctx, &taskReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "taskResp RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskResp))
}

func CreateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数失败"))
		return
	}
	taskRes, err := rpc.TaskCreate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "TaskList RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func GetTaskHandler(ctx *gin.Context) {

	var req pb.TaskRequest
	req.Id, _ = strconv.ParseUint(ctx.Param("id"), 10, 64)
	taskRes, err := rpc.TaskGet(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "TaskList RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func UpdateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数失败"))
		return
	}

	taskRes, err := rpc.TaskUpdate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "TaskUpdate RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func DeleteTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数失败"))
		return
	}
	taskRes, err := rpc.TaskDelete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "TaskDelete RPC 调度失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}
