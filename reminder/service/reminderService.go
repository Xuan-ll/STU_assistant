package service

import (
	"context"
	"sync"

	"stu_Assistant/idl/pb"
	"stu_Assistant/reminder/reminderconfig"
	"stu_Assistant/reminder/repository/cache"
	"stu_Assistant/reminder/repository/model"
	"stu_Assistant/reminder/repository/ormdb"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var ReminderSrvIns *ReminderSrv
var ReminderSrvOnce sync.Once

type ReminderSrv struct {
}

func GetReminderSrv() *ReminderSrv {
	ReminderSrvOnce.Do(func() {
		ReminderSrvIns = &ReminderSrv{}
	})
	return ReminderSrvIns
}

func (r *ReminderSrv) UpdateDeadLine(ctx context.Context, req *pb.ReminderRequest, resp *pb.ReminderListRespose) (err error) {
	resp.Code = reminderconfig.SUCCESS
	user_dealine, find_error := ormdb.NewReminderDao(ctx).FindReminder(uint(req.Uid))
	if find_error == gorm.ErrRecordNotFound {
		reminder := &model.Reminder{Uid: uint(req.Uid), Deadline: req.DeadLine}
		err = ormdb.NewReminderDao(ctx).CreateReminder(reminder)
		if err != nil {
			resp.Code = reminderconfig.ERROR
			return
		}
		err = cache.RemoveTasks(ctx, req.Uid)
		if err != nil {
			resp.Code = reminderconfig.ERROR
			return
		}
		var tasks []redis.Z
		tasks, err = cache.GetTasks(ctx, req.Uid, req.DeadLine)
		if err != nil {
			resp.Code = reminderconfig.ERROR
			return
		}
		var taskRes []*pb.ReminderModel
		resp.Count = uint32(len(tasks))
		if resp.Count != 0 {
			for _, item := range tasks {
				taskRes = append(taskRes, BuildReminder(item.Member.(string), int64(item.Score)))
			}	
		}	
		resp.TaskList = taskRes
	} else if find_error == nil {
		user_dealine.Deadline = req.DeadLine
		err = ormdb.NewReminderDao(ctx).UpdateReminder(user_dealine)
		if err != nil {
			resp.Code = reminderconfig.ERROR
			return
		}
		err = cache.RemoveTasks(ctx, req.Uid)
		if err != nil {
			resp.Code = reminderconfig.ERROR
			return
		}
		var tasks []redis.Z
		tasks, err = cache.GetTasks(ctx, req.Uid, req.DeadLine)
		if err != nil {
			resp.Code = reminderconfig.ERROR
			return
		}
		var taskRes []*pb.ReminderModel
		resp.Count = uint32(len(tasks))
		if resp.Count != 0 {
			for _, item := range tasks {
				taskRes = append(taskRes, BuildReminder(item.Member.(string), int64(item.Score)))
			}	
		}	
		resp.TaskList = taskRes		
	} else {
		resp.Code = reminderconfig.ERROR
		return
	}
	return
}

func (r *ReminderSrv) GetTaskInCache(ctx context.Context, req *pb.GetTaskRequest, resp *pb.ReminderListRespose) (err error) {
	resp.Code = reminderconfig.SUCCESS
	var user_deadline *model.Reminder
	user_deadline, err = ormdb.NewReminderDao(ctx).FindReminder(uint(req.Uid))
	if err != nil {
		resp.Code = reminderconfig.ERROR
		return
	}
	deadline := user_deadline.Deadline
	err = cache.RemoveTasks(ctx, req.Uid)
	if err != nil {
		resp.Code = reminderconfig.ERROR
		return
	}
	var tasks []redis.Z
	tasks, err = cache.GetTasks(ctx, req.Uid, deadline)
	if err != nil {
		resp.Code = reminderconfig.ERROR
		return
	}
    var taskRes []*pb.ReminderModel
	resp.Count = uint32(len(tasks))
	if resp.Count != 0 {
		for _, item := range tasks {
			taskRes = append(taskRes, BuildReminder(item.Member.(string), int64(item.Score)))
		}	
	}
	resp.TaskList = taskRes
    
	return
}

func BuildReminder(title string, deadline int64) *pb.ReminderModel {
	return &pb.ReminderModel{Title: title, EndTime: deadline}
}
