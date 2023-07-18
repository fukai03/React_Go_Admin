package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义相应函数，接收gin.Context、http状态码、自定义状态码、数据、消息作为参数
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	// 若data为nil，则返回空数组
	emptyArray := []string{}

	if data == nil {
		ctx.JSON(httpStatus, gin.H{"code": code, "data": gin.H{"result": emptyArray}, "msg": msg})
	} else {
		ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
	}
}

// 常用的响应函数

// ResponseOk 响应成功
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

// ResponseFail 响应失败
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
