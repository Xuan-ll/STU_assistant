package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Uid        uint   `gorm:"not null"`
	CourseName string `gorm:"not null"`
	StartClass uint   `gorm:"not null"`
	EndClass   uint   `gorm:"not null"`
	StartWeek  uint   `gorm:"not null"`
	EndWeek    uint   `gorm:"not null"`
	DayOfWeek  uint   `gorm:"not null"`
}
