package model

import (
	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Age       int
	Position  string
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt
}

func (User) TableName() string {
	return "user"
}
