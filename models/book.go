package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Title       string
	Author      string
	Cover       *[]byte `gorm:"type:longblob"`
	PublishedAt *time.Time
}
