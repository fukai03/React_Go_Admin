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

	categoryRoutes := r.Group("/categories")
	categoryRoutes.Use(middleware.AuthMiddleware(), middleware.RecoveryMiddleware()) // 使用中间件保护categories接口
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update) // put请求更新所有字段,相当于直接替换
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)
	// categoryRoutes.PATCH("/:id", categoryController.Update) // patch请求只更新部分字段
	return r
}
