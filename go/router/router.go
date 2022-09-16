package router

import (
	"gin_blogs_first/controller"

	"github.com/gin-gonic/gin"
)

//启动gin服务
func Start() {
	r := gin.Default()
	r.POST("/user", controller.RegisterUser)

	r.Run()
}
