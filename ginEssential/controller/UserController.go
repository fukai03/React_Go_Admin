package controller

import (
	"ginEssential/common"
	"ginEssential/dto"
	"ginEssential/model"
	"ginEssential/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {

	db := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}
	if name == "" {
		name = utils.RandomString(10)
	}
	// 判断手机号是否存在,需要查找数据库
	if isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已经存在",
		})
		return
	}
	// 创建用户(用户不存在)
	// 加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// 加密系统错误
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "加密错误",
		})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	db.Create(&newUser)

	// utils.tele()

	// 返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"user": name,
	})
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}
	// 判断手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}

	// 判断密码是否正确(输入的密码和数据库)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}

	// 验证通过，使用jwt包发放token
	token, err := common.ReleaseToken(user)
	// 创建token失败时
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		// 记录日志
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{"token": token},
	})
}

// Info 用户信息,此时用户已经通过了中间件的验证，可以从上下文中获取用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	})
}

// 判断手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	// if user.ID != 0 {
	// 	return true
	// }
	return user.ID != 0
}
