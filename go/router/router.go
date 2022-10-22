package router

import (
	"gin_blogs_first/controller"
	"gin_blogs_first/myCors"

	"github.com/gin-gonic/gin"
)

//启动gin服务
func Start() {
	r := gin.Default()

	// 自己写的跨域中间件
	r.Use(myCors.Cors())

	// gin自带的跨域中间件 (但是还是有问题 不能跨域)
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// r.Use(cors.New(config))
	// 等同于上面三行代码的写法
	// r.Use(cors.Default())

	r.POST("/user", controller.RegisterUser)
	r.POST("/login", controller.Login)

	r.GET("/editor/:pid", controller.GetEditor)
	r.POST("/editor", controller.AddEditor)
	r.GET("/editorAllList", controller.GetAllEditor)

	r.Run()
}
