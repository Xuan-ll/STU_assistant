package service

import (
	"context"
	"sync"

	"gorm.io/gorm"

	"stu_Assistant/idl/pb"
	"stu_Assistant/task/repository/model"
	"stu_Assistant/task/repository/ormdb"
	"stu_Assistant/task/repository/cache"
	"stu_Assistant/task/taskconfig"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (t *TaskSrv) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code  = taskconfig.SUCCESS
	_, err = ormdb.NewTaskDao(ctx).CheckTaskByTitle(req.Title)
	if err != nil {
		if err == gorm.ErrRecordNotFound { // 如果不存在就继续下去
			// ...continue
		} else {
			resp.Code = taskconfig.ERROR
			return
		}
	}	
	m := &model.Task{
		Uid:       uint(req.Uid),
		Title:     req.Title,
		Status:    int(req.Status),
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	err = cache.AddOrUpdateRedis(ctx, m.Uid, m.Title, m.Content, m.EndTime)
	if err!= nil {
		resp.Code = taskconfig.ERROR
		return
	}
	return ormdb.NewTaskDao(ctx).CreateTask(m)
}

// GetTasksList 实现备忘录服务接口 获取备忘录列表
func (t *TaskSrv) GetTasksList(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskListResponse) (err error) {
	resp.Code = taskconfig.SUCCESS
	if req.Limit == 0 {
		req.Limit = 10
	}
	// 查找备忘录
	r, count, err := ormdb.NewTaskDao(ctx).ListTaskByUserId(req.Uid, int(req.Start), int(req.Limit))
	if err != nil {
		resp.Code = taskconfig.ERROR
		return
	}
	// 返回proto里面定义的类型
	var taskRes []*pb.TaskModel
	for _, item := range r {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = uint32(count)   // task的总数
	return
}

// GetTask 获取详细的备忘录
func (t *TaskSrv) GetTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = taskconfig.SUCCESS
	r, err := ormdb.NewTaskDao(ctx).GetTaskByTaskIdAndUserId(req.Id, req.Uid)
	if err != nil {
		resp.Code = taskconfig.ERROR
		return
	}
	taskRes := BuildTask(r)
	resp.TaskDetail = taskRes
	return
}

// UpdateTask 修改备忘录
func (t *TaskSrv) UpdateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	// 查找该用户的这条信息
	resp.Code = taskconfig.SUCCESS
    oldTaskData, err:= ormdb.NewTaskDao(ctx).GetTaskByTaskIdAndUserId(req.Id, req.Uid)
	if err!= nil {
		resp.Code = taskconfig.ERROR
		return
	}
	taskData, err := ormdb.NewTaskDao(ctx).UpdateTask(req)
	if err != nil {
		resp.Code = taskconfig.ERROR
		return
	}
	err = cache.RemoveRedis(ctx, uint(oldTaskData.Uid), oldTaskData.Title, oldTaskData.Content)
	if err!= nil {
		resp.Code = taskconfig.ERROR
		return
	}
    if req.EndTime == 0{
		req.EndTime = oldTaskData.EndTime
	}
	err = cache.AddOrUpdateRedis(ctx, uint(req.Uid), req.Title, req.Content, req.EndTime) //!!!!
	if err!= nil {
		resp.Code = taskconfig.ERROR
		return
	}
	resp.TaskDetail = BuildTask(taskData)
	return
}

// DeleteTask 删除备忘录
func (t *TaskSrv) DeleteTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = taskconfig.SUCCESS
	err = ormdb.NewTaskDao(ctx).DeleteTaskByIdAndUserId(req.Id, req.Uid)
	if err != nil {
		resp.Code = taskconfig.ERROR
		return
	}
    err = cache.RemoveRedis(ctx, uint(req.Uid), req.Title, req.Content)  
	return
}

func BuildTask(item *model.Task) *pb.TaskModel {
	taskModel := pb.TaskModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
	return &taskModel
}