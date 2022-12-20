package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"mes/internal/database"
)

// Text 使用Multipart表单,上传字符串
func Text(c *gin.Context) {
	// 获取传递来的比特流数据
	rawData, err := c.GetRawData()
	if err != nil {
		log.Panicln(err)
	}

	text := database.Text{
		Model:   gorm.Model{},
		Content: string(rawData),
	}

	db, err := database.NewSqliteDB()
	if err != nil {
		log.Panicln(err)
	}

	err = db.AutoMigrate(&database.Text{})
	if err != nil {
		log.Panicln(err)
	}

	db.Create(&text)

	fmt.Println(string(rawData))
}
