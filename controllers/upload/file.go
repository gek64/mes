package upload

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

var (
	DefaultDestinationKey = "destination"
	DefaultFileKey        = "fileName"
	DefaultDestination    = "upload"
	DefaultPerm           = os.FileMode(0755)
)

type Upload struct {
	FileName    string `form:"fileName" json:"fileName" xml:"fileName" uri:"fileName"`
	Destination string `form:"destination" json:"destination" xml:"destination" uri:"destination"`
}

// Raw 以 application/octet-stream 方式上传单个文件
func Raw(c *gin.Context) {
	// 获取url中的目标文件名
	var u Upload
	err := c.ShouldBindUri(&u)
	if err != nil {
		log.Panicln(err)
	}

	// 获取传递来的比特流数据
	rawData, err := c.GetRawData()
	if err != nil {
		log.Panicln(err)
	}

	// 新建文件存储目的地
	p, err := filepath.Abs(DefaultDestination)
	if err != nil {
		log.Panicln(err)
	}
	err = os.MkdirAll(p, DefaultPerm)
	if err != nil {
		log.Panicln(err)
	}
	// 新建文件
	f := filepath.Join(p, u.FileName)
	tmpFile, err := os.Create(f)
	// 函数结束时关闭文件
	defer func(tmpFile *os.File) {
		err := tmpFile.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(tmpFile)

	// 使用缓冲写入器
	writer := bufio.NewWriter(tmpFile)
	// 文件中写入比特流
	_, err = writer.Write(rawData)
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
	location, err := filepath.Abs(c.DefaultPostForm(DefaultDestinationKey, ""))
	if err != nil {
		log.Panicln(err)
	}
	// 存储路径不存在则新建
	_, err = os.Stat(location)
	if os.IsNotExist(err) {
		err := os.MkdirAll(location, DefaultPerm)
		if err != nil {
			log.Panicln(err)
		}
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
