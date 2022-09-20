package routers

import (
	"github.com/gin-gonic/gin"
	"mes/controllers/upload"
)

func Upload(router *gin.Engine) {
	routerGroup := router.Group("/upload")
	{
		routerGroup.POST("/raw/:fileName", upload.Raw)
		routerGroup.POST("/files", upload.Files)
	}
}
