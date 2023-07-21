package main

import (
	"ginEssential/common"

	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
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
	// 初始化配置
	InitConfig()

	// 初始化数据库
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" { // 若配置文件中有端口号，则使用配置文件中的端口号
		panic(r.Run(":" + port))
	}
	r.Run() // 否则使用默认端口号8080
}

func InitConfig() {
	// 获取路径
	workDir, _ := os.Getwd()
	// viper配置
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
