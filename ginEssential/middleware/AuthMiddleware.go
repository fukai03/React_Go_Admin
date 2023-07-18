package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 验证token格式,若为空或者不是以Bearer开头则认为是无效token
		if tokenString == "" || tokenString[:7] != "Bearer " {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		// 截取token字符串
		tokenString = tokenString[7:]
		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid { // 解析失败或者token无效
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		// 验证通过后获取claims中的userId
		userId := claims.UserId
		Db := common.GetDB()
		var user model.User
		Db.First(&user, userId) // 根据userId查找用户

		// 用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		// 用户存在，将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
