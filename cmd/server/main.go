package main

import (
	"github.com/gin-gonic/gin"
	"log"
	routers2 "mes/internal/server/routers"
)

func main() {
	// 运行模式
	gin.SetMode(gin.DebugMode)

	// 创建默认路由引擎
	engine := gin.Default()

	// 加载路由
	routers2.File(engine)
	routers2.Text(engine)

	// 在:80上启动
	err := engine.Run("127.0.0.1:80")
	if err != nil {
		log.Panicln(err)
	}
}
