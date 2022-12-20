package database

import (
	"gorm.io/gorm"
)

type Text struct {
	gorm.Model
	Content string
}
