package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICategoryController interface {
	RestController
}

// 为了controller层的代码更加清晰，我们将Category相关的代码放到CategoryController.go中
type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()              // 获取数据库连接
	db.AutoMigrate(&model.Category{}) // 自动迁移
	return CategoryController{DB: db}
}

// Create 创建分类
func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category // 接收请求数据
	// name := ctx.PostForm("name")       // 获取参数
	// fmt.Println("name", name)
	ctx.BindJSON(&requestCategory) // 绑定请求数据,不能使用Bind，因为Bind无法接收Json数据

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
	} else {
		c.DB.Create(&requestCategory) // 创建分类
		// q: 为什么数据库中时间显示正常，返回的时间却是0？
		// a: 因为我们在model/Category.go中定义了json:"created_at"，所以返回的时间是0
		// q: 怎么解决？
		// a: 在model/Category.go中删除json:"created_at"，或者在model/Category.go中定义json:"created_at"，在model/Category.go中定义gorm:"autoCreateTime"，这样就可以了
		response.Success(ctx, gin.H{"category": requestCategory}, "创建成功")

	}
}

// Update 更新分类
func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory model.Category
	ctx.BindJSON(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
	}
	// 获取path中的参数,categoryID是string类型，需要转换成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category

	if c.DB.First(&updateCategory, categoryId).Error != nil {
		response.Fail(ctx, nil, "分类不存在")
	}

	// 更新分类
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "更新成功")
}

// Show 查看分类
func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数,categoryID是string类型，需要转换成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category model.Category

	if c.DB.First(&category, categoryId).Error != nil {
		response.Fail(ctx, nil, "分类不存在")
	} else {
		response.Success(ctx, gin.H{"category": category}, "查询成功")
	}
}

// Delete 删除分类
func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数,categoryID是string类型，需要转换成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// err := c.DB.Delete(model.Category{}, categoryId).Error
	// 查询数据库是否有id
	var category model.Category
	if c.DB.First(&category, categoryId).Error != nil {
		response.Fail(ctx, nil, "删除失败，分类不存在")
		return
	}
	// 删除分类
	err := c.DB.Delete(model.Category{}, categoryId).Error
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}
	response.Success(ctx, nil, "删除成功")
}
