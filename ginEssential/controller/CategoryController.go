package controller

import (
	"ginEssential/model"
	"ginEssential/repository"
	"ginEssential/response"
	"ginEssential/vo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}

// 为了controller层的代码更加清晰，我们将Category相关的代码放到CategoryController.go中
type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(&model.Category{}) // 自动迁移
	return CategoryController{Repository: repository}
}

// Create 创建分类
func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest // 接收请求数据
	err := ctx.ShouldBindJSON(&requestCategory)
	if err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}
	category, err := c.Repository.Create(requestCategory.Name) // 创建分类
	if err != nil {
		panic(err)
	}
	c.Repository.Create(requestCategory.Name) // 创建分类
	response.Success(ctx, gin.H{"category": category}, "创建成功")

}

// Update 更新分类
func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory vo.CreateCategoryRequest // 接收请求数据
	err := ctx.ShouldBindJSON(&requestCategory)
	if err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}
	// 获取path中的参数,categoryID是string类型，需要转换成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectById(categoryId) // 查询分类是否存在
	if err != nil {
		panic(err)
	}
	// 更新分类
	c.Repository.Update(*updateCategory, requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "更新成功")
}

// Show 查看分类
func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数,categoryID是string类型，需要转换成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId) // 查询分类是否存在
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "查询成功")
}

// Delete 删除分类
func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	err := c.Repository.DeleteById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "删除失败，分类不存在")
		return
	}
	response.Success(ctx, nil, "删除成功")
}
