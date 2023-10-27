package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Title       string     `gorm:"size:500"`
	Author      string     `gorm:"size:500"`
	Cover       *[]byte    `gorm:"type:longblob"`
	PublishedAt *time.Time // maybe we need to store only date here.
}
