package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/vo"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IPostController interface {
	RestController
	ListController
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(&model.Post{})
	return PostController{DB: db}
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	err := ctx.ShouldBindJSON(&requestPost)
	if err != nil {
		panic(err)
	}
	// 获取用户信息，前面使用中间件将用户信息保存到了上下文中
	user, _ := ctx.Get("user")

	// 创建post
	post := model.Post{
		UserID:     user.(model.User).ID,
		CategoryID: requestPost.CategoryID,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	// 数据库创建post
	err = p.DB.Create(&post).Error
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"id": post.ID}, "创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	err := ctx.ShouldBindJSON(&requestPost)
	if err != nil {
		panic(err)
	}

	user, _ := ctx.Get("user")

	postId := ctx.Params.ByName("id")

	var post model.Post
	// 此处需要使用Preload("Category")，否则返回的post中的Category为空
	postErr := p.DB.Preload("Category").Where("id = ?", postId).First(&post).Error
	if postErr != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	if post.UserID != user.(model.User).ID {
		response.Fail(ctx, nil, "文章不属于你，无法修改")
		return
	}

	// 更新文章
	if err = p.DB.Model(&post).Updates(requestPost).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "更新成功")
}

func (p PostController) Show(ctx *gin.Context) {
	postId := ctx.Params.ByName("id")

	var post model.Post
	if err := p.DB.Preload("Category").Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "成功")
}

func (p PostController) Delete(ctx *gin.Context) {
	postId := ctx.Params.ByName("id")

	var post model.Post
	if err := p.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	user, _ := ctx.Get("user")

	if post.UserID != user.(model.User).ID {
		response.Fail(ctx, nil, "文章不属于你，无法删除")
		return
	}

	if err := p.DB.Where("id = ?", postId).Delete(&post).Error; err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "删除成功")
}

func (p PostController) List(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	// 分页
	var posts []model.Post
	p.DB.Preload("Category").Order("create_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 获取总条数
	var total int64
	p.DB.Model(model.Post{}).Count(&total)

	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")

}
