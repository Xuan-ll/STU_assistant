package dao

import (
	"stu_Assistant/course/repository/db/model"
)

func migration() {
	_db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&model.Course{})
}
