package dao

import (
	"context"

	"gorm.io/gorm"

	"stu_Assistant/course/repository/db/model"
	"stu_Assistant/idl/pb"
)

type CourseDao struct {
	*gorm.DB
}

func NewCourseDao(ctx context.Context) *CourseDao {
	return &CourseDao{NewDBClient(ctx)}
}

// ListCourseByUserId 获取user id
func (dao *CourseDao) ListCourseByUserId(userId uint64, week uint32) (r []*model.Course, count int64, err error) {
	err = dao.Model(&model.Course{}).Where("uid = ?", userId).
		Where("start_week <=? AND? <= end_week", week, week).
		Find(&r).Count(&count).Error
	return
}

// UpdateCourse 更新course
func (dao *CourseDao) UpdateCourse(req *pb.CourseRequest) (r *model.Course, err error) {
	r = new(model.Course)
	err = dao.Model(&model.Course{}).
		Where("id = ? AND uid = ?", req.Id, req.Uid).
		First(&r).Error
	if err != nil {
		return
	}
	r.CourseName = req.CourseName
	r.StartClass = uint(req.StartClass)
	r.EndClass = uint(req.EndClass)
	r.StartWeek = uint(req.StartWeek)
	r.EndWeek = uint(req.EndWeek)
	r.DayOfWeek = uint(req.DayOfWeek)

	err = dao.Save(&r).Error
	return
}

func (dao *CourseDao) DeleteCourseByIdAndUserId(id, uId uint64) error {
	return dao.Model(&model.Course{}).
		Where("id =? AND uid=?", id, uId).
		Delete(&model.Course{}).Error
}

func (dao *CourseDao) CreateCourse(in *model.Course) error {
	return dao.Model(&model.Course{}).Create(&in).Error
}
