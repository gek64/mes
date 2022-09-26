package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mes/routers"
)

func main() {
	// 创建默认路由引擎
	engine := gin.Default()

	// 加载路由
	routers.Upload(engine)

	// 在:80上启动
	err := engine.Run("localhost:80")
	if err != nil {
		log.Panicln(err)
	}
}
