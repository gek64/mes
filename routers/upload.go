package routers

import (
	"github.com/gin-gonic/gin"
	"mes/controllers/upload"
)

func Upload(router *gin.Engine) {
	routerGroup := router.Group("/upload")
	{
		// 文件
		routerGroup.POST("/rawFile", upload.RawFile)
		routerGroup.POST("/files", upload.Files)
		// 文本
		routerGroup.POST("/text", upload.Text)
	}
}
