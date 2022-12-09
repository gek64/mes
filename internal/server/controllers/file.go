package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"mes/pkg/mod"
	"path/filepath"
)

var (
	DefaultDestinationKey = "destination"
	DefaultFileKey        = "file"
	DefaultDestination    = "upload"
)

type File struct {
	FileName    string `form:"fileName" json:"fileName" xml:"fileName" uri:"fileName"`
	Destination string `form:"destination" json:"destination" xml:"destination" uri:"destination"`
}

// RawFile 以 application/octet-stream 方式上传单个文件
func RawFile(c *gin.Context) {
	// url变量绑定到变量,获取url中的文件名
	var fileDetail File
	err := c.ShouldBindQuery(&fileDetail)
	if err != nil {
		log.Panicln(err)
	}
	if fileDetail.Destination == "" {
		fileDetail.Destination = DefaultDestination
	}

	// 获取传递来的比特流数据
	rawData, err := c.GetRawData()
	if err != nil {
		log.Panicln(err)
	}

	// 新建文件存储目的地
	err = mod.CreateFolderIFNotExist(fileDetail.Destination)
	if err != nil {
		log.Panicln(err)
	}

	// 新建文件
	err = mod.WriteToFile(filepath.Join(fileDetail.Destination, fileDetail.FileName), rawData)
	if err != nil {
		log.Panicln(err)
	}
}

// Files 使用Multipart表单,上传一个或者多个文件
func Files(c *gin.Context) {
	// 获取请求中的所有Multipart表单
	form, err := c.MultipartForm()
	if err != nil {
		log.Panicln(err)
	}

	// 获取表单中的上传文件的存储路径
	location, err := filepath.Abs(c.DefaultPostForm(DefaultDestinationKey, DefaultDestination))
	if err != nil {
		log.Panicln(err)
	}
	// 存储路径不存在则新建
	err = mod.CreateFolderIFNotExist(location)
	if err != nil {
		log.Panicln(err)
	}

	// 从表单中获取所有value类型为File的项中name为"file"的所有项
	files := form.File[DefaultFileKey]

	for _, file := range files {
		// 文件保存到指定位置,文件夹位置要事先建立,不然会报错
		err := c.SaveUploadedFile(file, filepath.Join(location, file.Filename))
		if err != nil {
			log.Panicln(err)
		}
	}
}
