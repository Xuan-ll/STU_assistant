package model

import (
	"gorm.io/gorm"
)
type Task struct {
	gorm.Model
	Uid          uint    `gorm:"not null"`
	Title        string  `gorm:"index; not null"`
	Status       int     `gorm:"default:0"`
    Content      string  `gorm:"type:longtext"`
	StartTime    int64
	EndTime      int64   `gorm:"default:0"`
}// 建立数据库时会自动加载id字段并作为主键