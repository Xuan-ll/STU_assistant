package rpc

import (
	"context"

	"stu_Assistant/idl/pb"
	"stu_Assistant/gateway/gwconfig"
)	

func ReminderUpdate(ctx context.Context, req *pb.ReminderRequest) (resp *pb.ReminderListRespose, err error) {
	r, err := ReminderService.UpdateDeadLine(ctx, req)
	if err!= nil {
		return
	}
	if r.Code!= gwconfig.SUCCESS {
		return
	}

	return r, nil
}

func ReminderGet(ctx context.Context, req *pb.GetTaskRequest) (resp *pb.ReminderListRespose, err error) {
	r, err := ReminderService.GetTaskInCache(ctx, req)
	if err!= nil {
		return
	}
	if r.Code!= gwconfig.SUCCESS {
		return
	}

	return r, nil
}