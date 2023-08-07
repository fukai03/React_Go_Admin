package controller

import (
	"github.com/gin-gonic/gin"
)

// 定义通用接口,包含增删改查,方便复用
type RestController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
