package router

import (
	"github.com/gin-gonic/gin"
	"go-gin/app/controller"
	"go-gin/app/controller/api"
	"go-gin/app/middleware"
)

func SetupRouter() *gin.Engine {
	//router := gin.Default()
	router := gin.New()
	router.Use(middleware.Logger(), gin.Recovery())
	// 添加 Get 请求路由
	router.GET("/", controller.Index)
	router.GET("/welcome", controller.Welcome)
	v1 := router.Group("/")
	{
		v1.POST("/user/login", new(api.User).Login)
		v1.POST("/user/register", new(api.User).Register)
		v1.GET("/user/info", new(api.User).UserInfo)
	}
	v2 := router.Group("admin")
	{
		v2.POST("/login", new(api.User).Login)
	}
	return router
}