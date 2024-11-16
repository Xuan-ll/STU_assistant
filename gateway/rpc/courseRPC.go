package rpc

import (
	"context"

	"stu_Assistant/gateway/gwconfig"
	"stu_Assistant/idl/pb"
)

func CourseCreate(ctx context.Context, req *pb.CourseRequest) (resp *pb.CourseDetailResponse, err error) {
	r, err := CourseService.CreateCourse(ctx, req)
	if err != nil {
		return
	}
	if r.Code != gwconfig.SUCCESS {
		return
	}

	return r, nil
}

func CourseUpdate(ctx context.Context, req *pb.CourseRequest) (resp *pb.CourseDetailResponse, err error) {
	r, err := CourseService.UpdateCourse(ctx, req)
	if err != nil {
		return
	}
	if r.Code != gwconfig.SUCCESS {
		return
	}

	return r, nil
}

func CourseDelete(ctx context.Context, req *pb.CourseRequest) (resp *pb.CourseDetailResponse, err error) {
	r, err := CourseService.DeleteCourse(ctx, req)
	if err != nil {
		return
	}
	if r.Code != gwconfig.SUCCESS {
		return
	}

	return r, nil
}

func CourseList(ctx context.Context, req *pb.CourseRequest) (resp *pb.CourseListResponse, err error) {
	r, err := CourseService.GetCoursesList(ctx, req)
	if err != nil {
		return
	}
	if r.Code != gwconfig.SUCCESS {
		return
	}

	return r, nil
}
