package routers

import (
	"github.com/gin-gonic/gin"
	"mes/controllers/upload"
)

func Upload(router *gin.Engine) {
	routerGroup := router.Group("/upload")
	{
		routerGroup.POST("/raw", upload.UpdateRawData)
		routerGroup.POST("/", upload.UpdateFiles)
	}
}
