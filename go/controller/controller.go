package controller

import (
	"gin_blogs_first/dao"
	"gin_blogs_first/model"

	"github.com/gin-gonic/gin"
)

//注册用户
func RegisterUser(c *gin.Context) {
	user := model.User{}
	c.BindJSON(&user)

	b := dao.Mgr.RegisterUser(&user)
	if b {
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "非法输入或者用户已存在",
		})
	}
}

//登录
func Login(c *gin.Context) {
	user := model.User{}
	c.BindJSON(&user)
	u := dao.Mgr.Login(user.Username)

	if u.Username == "" {
		c.JSON(400, gin.H{
			"message": "无此用户",
		})
	} else {
		c.JSON(200, gin.H{
			"message": user.Username + "欢迎登录",
		})
	}
}
