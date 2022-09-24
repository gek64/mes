package routers

import (
	"github.com/gin-gonic/gin"
	"mes/controllers/upload"
)

func Upload(router *gin.Engine) {
	routerGroup := router.Group("/upload")
	{
		routerGroup.POST("/raw", upload.Raw)
		routerGroup.POST("/file", upload.Files)
	}
}
