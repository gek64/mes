package routers

import (
	"github.com/gin-gonic/gin"
	"mes/internal/server/controllers"
)

func Text(router *gin.Engine) {
	routerGroup := router.Group("/text")
	{
		routerGroup.POST("/", controllers.Text)
	}
}
