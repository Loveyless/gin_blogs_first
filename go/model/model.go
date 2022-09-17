package model

import (
	"gorm.io/gorm"
)

//用户模型
type User struct {
	gorm.Model
	// Username string `json:"username" form:"username"`
	// Password string `json:"password" form:"password"`
	Username string `json:"username"`
	Password string `json:"password"`
}
