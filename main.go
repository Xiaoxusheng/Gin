package main

import (
	"Gin/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := router.Router()
	//开启日志颜色
	gin.ForceConsoleColor()
	//静态文件
	r.Static("/img", "./img")
	//允许跨域
	r.Use(cors.Default())

	r.Run(":8080") // 监听并在 0.0.0.0:80 上启动服务
}
