package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"rpcoptool/pkg/utils"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code uint32

		code = http.StatusOK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			code = http.StatusUnauthorized
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			code = http.StatusUnauthorized
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "权限过期，请重新登陆",
			})
			c.Abort()
			return
		}
		//用户登陆获取token,然后使用token进去鉴权网址,通过鉴权则缓存用户信息,维持一个登陆状态

		c.Next()
	}
}
