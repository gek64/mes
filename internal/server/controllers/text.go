package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// Text 使用Multipart表单,上传字符串
func Text(c *gin.Context) {
	// 获取传递来的比特流数据
	rawData, err := c.GetRawData()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(rawData))
}
