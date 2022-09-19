package model

import (
	"gorm.io/gorm"
)

//用户模型
type User struct {
	gorm.Model
	// Username string `json:"username" form:"username"`
	// Password string `json:"password" form:"password"`
	Username string `json:"username" binding:"required,min=3,max=10"`
	Password string `json:"password" binding:"required,min=3,max=15"`
}

//博客模型
type Editor struct {
	gorm.Model
	UserId  string `json:"user_id"  binding:"required"`
	Content string `gorm:"type:text"  binding:"required"`
	Tag     string
}
