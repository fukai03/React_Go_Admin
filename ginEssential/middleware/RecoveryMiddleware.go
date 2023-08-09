package middleware

import (
	"fmt"
	"ginEssential/response"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 捕获panic异常的中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, gin.H{"err": fmt.Sprint(err)}, "系统异常")
			}
		}()

		ctx.Next()
	}
}
