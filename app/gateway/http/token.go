package http

import (
	"net/http"
	log "rpcoptool/pkg/logger"
	"rpcoptool/pkg/utils"

	"rpcoptool/pkg/ctl"

	"github.com/gin-gonic/gin"
)

type TokenGet struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

func GetToken(ctx *gin.Context) {
	var req TokenGet
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "GetToken Bind 绑定参数失败"))
		return
	}
	token, err := utils.GenerateToken(uint(req.Id))
	if err != nil {
		log.LogrusObj.Errorf("Generate Token fail %v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	res := &TokenData{
		User:  req.Name,
		Token: token,
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, res))
}
