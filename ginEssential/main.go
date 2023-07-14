package main

import (
	"ginEssential/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//	func main() {
//		r := gin.Default()
//		r.GET("/ping", func(c *gin.Context) {
//			c.JSON(http.StatusOK, gin.H{
//				"message": "pong",
//			})
//		})
//		r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//	}

func main() {
	// 连接数据库
	// db := InitDB()
	// defer db.Close()

	// 初始化数据库
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	r.Run(":8888")
}
