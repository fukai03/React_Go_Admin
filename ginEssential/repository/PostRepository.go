package repository

import (
	"ginEssential/common"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository() PostRepository {
	return PostRepository{DB: common.GetDB()}
}
