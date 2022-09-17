package controller

import (
	"fmt"
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

	// buf := make([]byte, 2048)
	// n, _ := c.Request.Body.Read(buf)
	// fmt.Println("原始请求数据", string(buf[0:n]))

	s := c.Request.Header.Get("Content-Type")
	fmt.Println(s)

	user := model.User{}
	c.ShouldBind(&user)

	fmt.Println("拿到的请求数据")
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	//查这个用户名
	u := dao.Mgr.Login(user.Username)
	// 对比
	if u.Username == "" {
		c.JSON(400, gin.H{
			"message": "无此用户",
		})
	} else {
		if u.Password == user.Password {
			c.JSON(200, gin.H{
				"message": user.Username + "欢迎登录",
			})
		} else {
			c.JSON(400, gin.H{
				"message": "用户名或密码错误",
			})
		}
	}
}
