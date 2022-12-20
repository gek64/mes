package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mes/internal/routers"
)

func main() {
	// 运行模式
	gin.SetMode(gin.DebugMode)

	// 创建默认路由引擎
	engine := gin.Default()

	// 加载路由
	routers.File(engine)
	routers.Text(engine)

	// 在:80上启动
	err := engine.Run("127.0.0.1:80")
	if err != nil {
		log.Panicln(err)
	}
}
