package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteDB() (db *gorm.DB, err error) {
	// 自定义sqlite数据库文件
	s := sqlite.Open("mes.db")

	// 返回数据库
	return gorm.Open(s, &gorm.Config{})
}
