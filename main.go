package main

import (
	"gin-project-base/middleware"
	"gin-project-base/router"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main(){
	app := gin.New()

	//设置日志写入到文件
	f, _  := os.OpenFile("gin.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(f)

	//打印调试
	log.SetPrefix("debug:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("this is a test message!")

	//设置中间件
	app.Use(gin.Recovery())
	app.Use(gin.LoggerWithFormatter(middleware.SetLogFormat))

	//设置静态资源路径
	app.Static("image", "./static/image")
	//设置HTML模板目录
	app.LoadHTMLGlob("view/*")
	//设置路由
	router.InitRouter(app)

	//启动app
	app.Run(":8080")

}