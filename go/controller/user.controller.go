package controller

import (
	"gin_blogs_first/dao"
	"gin_blogs_first/model"
	"gin_blogs_first/validator"

	"github.com/gin-gonic/gin"
)

//注册用户
func RegisterUser(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": validator.Translate(err),
			"data":    "",
		})
		return
	}

	b := dao.Mgr.RegisterUser(&user)
	if b {
		c.JSON(200, gin.H{
			"message": "注册成功",
			"status":  200,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "非法输入或者用户已存在",
			"status":  400,
		})
	}
}

//登录
func Login(c *gin.Context) {

	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": validator.Translate(err),
			"data":    "",
		})
		return
	}

	//查这个用户名
	u := dao.Mgr.Login(user.Username)
	// 对比
	if u.Username == "" {
		c.JSON(400, gin.H{
			"message": "无此用户",
			"status":  400,
		})
	} else {
		if u.Password == user.Password {
			c.JSON(200, gin.H{
				"message": user.Username + "欢迎登录",
				"status":  200,
				"data":    u,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "用户名或密码错误",
				"status":  400,
			})
		}
	}
}
