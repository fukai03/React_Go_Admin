package main

import (
	"ginEssential/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)                     // 注册
	r.POST("/api/auth/login", controller.Login)                           // 登录
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) // 获取用户信息,使用中间件AuthMiddleware保护用户信息接口
	return r
}
