package ormdb

import (
	"context"
	"stu_Assistant/reminder/repository/model"

	"gorm.io/gorm"
)

type ReminderDao struct {
	*gorm.DB
}

func NewReminderDao(ctx context.Context) *ReminderDao {
	if ctx == nil {
		ctx = context.Background()
	}
	goDB := _db.WithContext(ctx)
	return &ReminderDao{DB: goDB}
}	

func (dao *ReminderDao) CreateReminder(in *model.Reminder) (err error) {
	return dao.Model(&model.Reminder{}).Create(&in).Error
}

func (dao *ReminderDao) FindReminder(uid uint) (r *model.Reminder, err error) {
	err = dao.Model(&model.Reminder{}).Where("user_id = ?", uid).Find(&r).Error
	if err!= nil {
		return nil, err
    }
	return
}

func (dao *ReminderDao) UpdateReminder(in *model.Reminder) (err error) {
	return dao.Model(&model.Reminder{}).Where("user_id = ?", in.Uid).Updates(&in).Error
}