package mods

import (
	"log"
	"os"
	"path/filepath"
)

var (
	DefaultPerm = os.FileMode(0755)
)

// CreateFolderIFNotExist 创建目录,如果目录路径不存在则创建
func CreateFolderIFNotExist(path string) (err error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(absPath, DefaultPerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateFileIFNotExist 创建文件,如果文件路径不存在则创建
func CreateFileIFNotExist(path string) (file *os.File, err error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	return os.OpenFile(absPath, os.O_RDWR|os.O_CREATE, DefaultPerm)
}

// WriteToFile 向文件中写入数据([]byte)
func WriteToFile(path string, data []byte) (err error) {
	file, err := CreateFileIFNotExist(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(file)

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
