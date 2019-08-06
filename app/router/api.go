package router

import (
	"github.com/gin-gonic/gin"
	"go-gin/app/controller"
	"go-gin/app/controller/api"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 添加 Get 请求路由
	router.GET("/", controller.Index)
	router.GET("/welcome", controller.Welcome)
	router.POST("/api/login", api.Login)
	router.POST("/api/register", api.Register)
	router.GET("/api/userInfo/:id", api.UserInfo)
	return router
}