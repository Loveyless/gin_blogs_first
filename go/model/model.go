package model

import (
	"gorm.io/gorm"
)

//用户模型
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
