package ormdb

import (
	"context"

	"gorm.io/gorm"

	"stu_Assistant/task/repository/model"
	"stu_Assistant/idl/pb"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{_db.WithContext(ctx)}
}

// ListTaskByUserId 获取user id
func (dao *TaskDao) ListTaskByUserId(userId uint64, start, limit int) (r []*model.Task, count int64, err error) {
	err = dao.Model(&model.Task{}).Offset(start).
		Limit(limit).Where("uid = ?", userId).
		Find(&r).Error
    if err!= nil {
		return nil, 0, err
	}
	err = dao.Model(&model.Task{}).Where("uid = ?", userId).
		Count(&count).Error
	return
}

// GetTaskByTaskIdAndUserId 通过 task id 和 user id 获取task
func (dao *TaskDao) GetTaskByTaskIdAndUserId(taskId, userId uint64) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).
		Where("id = ? AND uid = ?", taskId, userId).
		First(&r).Error

	return
}

// UpdateTask 更新task
func (dao *TaskDao) UpdateTask(req *pb.TaskRequest) (r *model.Task, err error) {
	r = new(model.Task)
	err = dao.Model(&model.Task{}).
		Where("id = ? AND uid = ?", req.Id, req.Uid).
		First(&r).Error
	if err != nil {
		return
	}
	r.Title = req.Title
	r.Status = int(req.Status)
	r.Content = req.Content

	err = dao.Save(&r).Error
	return
}

func (dao *TaskDao) DeleteTaskByIdAndUserId(id, uId uint64) error {
	return dao.Model(&model.Task{}).
		Where("id =? AND uid=?", id, uId).
		Delete(&model.Task{}).Error

}

func (dao *TaskDao) CreateTask(in *model.Task) error {
	return dao.Model(&model.Task{}).Create(&in).Error
}

func (dao *TaskDao) CheckTaskByTitle(title string) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).
		Where("title =?", title).
		Find(&r).Error
	return
}