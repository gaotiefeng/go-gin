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
	router.POST("/api/login", api.Login)
	router.POST("/api/register", new(api.User).Register)
	router.GET("/api/userInfo/:id", new(api.User).UserInfo)
	return router
}