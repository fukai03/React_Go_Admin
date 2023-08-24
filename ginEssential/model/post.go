package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"` // 定长字符串类型
	UserID     uint      `json:"user_id" gorm:"not null"`
	CategoryID uint      `json:"category_id" gorm:"not null;"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(100);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
	CreateAt   Time   `json:"create_at" gorm:"autoCreateTime;type:timestamp"`
	UpdateAt   Time   `json:"update_at" gorm:"autoUpdateTime;type:timestamp"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	post.ID = uuid.NewV4()
	return nil
}
