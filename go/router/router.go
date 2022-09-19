package router

import (
	"gin_blogs_first/controller"
	"gin_blogs_first/cors"

	"github.com/gin-gonic/gin"
)

//启动gin服务
func Start() {
	r := gin.Default()

	r.Use(cors.Cors())

	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// r.Use(cors.New(config))

	// r.Use(cors.Default())

	r.POST("/user", controller.RegisterUser)
	r.POST("/login", controller.Login)

	r.GET("/editor/:pid", controller.GetEditor)
	r.POST("/editor", controller.AddEditor)
	r.GET("/editorAllList", controller.GetAllEditor)

	r.Run()
}
