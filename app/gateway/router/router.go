package router

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"rpcoptool/app/gateway/http"
	"rpcoptool/app/gateway/middleware"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})

		// 用户服务
		v1.POST("/token/get", http.GetToken)

		// 需要登录保护的TASK服务-增删改查
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("tasks", http.ListTaskHandler)
			authed.POST("task", http.CreateTaskHandler)
			authed.GET("task/:id", http.GetTaskHandler)       // task_id
			authed.PUT("task/:id", http.UpdateTaskHandler)    // task_id
			authed.DELETE("task/:id", http.DeleteTaskHandler) // task_id
		}
	}
	return ginRouter
}
