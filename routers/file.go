package routers

import (
	"github.com/gin-gonic/gin"
	"mes/controllers"
)

func File(router *gin.Engine) {
	routerGroup := router.Group("/file")
	{
		routerGroup.POST("/", controllers.Files)
		routerGroup.POST("/raw", controllers.RawFile)
	}
}
