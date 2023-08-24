package model

type Category struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"` // 设置主键并自增
	Name     string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreateAt Time   `json:"create_at" gorm:"autoCreateTime;type:timestamp"`
	UpdateAt Time   `json:"update_at" gorm:"autoUpdateTime;type:timestamp"`
}
