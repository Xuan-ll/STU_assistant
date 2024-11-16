package model

import (
	"gorm.io/gorm"
)

type Reminder struct {
	gorm.Model
	Uid          uint    `gorm:"index; not null; column:user_id"`
    Deadline     int64   `gorm:"default:259200"`  // 默认为3天
}