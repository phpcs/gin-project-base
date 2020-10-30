package router

import (
	"gin-project-base/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine){
	//一个路由
	router.POST("/upload", controller.Upload)

	//分组
	group1 := router.Group("/v1")
	{
		group1.GET("/index", controller.Index)
		group1.POST("/login", controller.Login)
		group1.GET("/user/:id", controller.GetUser)
	}
}