package service

import (
	"context"
	"sync"

	"stu_Assistant/course/config"
	"stu_Assistant/course/repository/db/dao"
	"stu_Assistant/course/repository/db/model"
	"stu_Assistant/idl/pb"
	// log "github.com/CocaineCong/micro-todoList/pkg/logger"
)

var CourseSrvIns *CourseSrv
var CourseSrvOnce sync.Once

type CourseSrv struct {
}

func GetCourseSrv() *CourseSrv {
	CourseSrvOnce.Do(func() {
		CourseSrvIns = &CourseSrv{}
	})
	return CourseSrvIns
}

// CreateCourse 创建课程，将课程信息生产
func (t *CourseSrv) CreateCourse(ctx context.Context, req *pb.CourseRequest, resp *pb.CourseDetailResponse) (err error) {
	resp.Code = config.SUCCESS
	m := &model.Course{
		Uid:        uint(req.Uid),
		CourseName: req.CourseName,
		StartClass: uint(req.StartClass),
		EndClass:   uint(req.EndClass),
		StartWeek:  uint(req.StartWeek),
		EndWeek:    uint(req.EndWeek),
		DayOfWeek:  uint(req.DayOfWeek),
	}
	return dao.NewCourseDao(ctx).CreateCourse(m)
}
func (t *CourseSrv) GetCoursesList(ctx context.Context, req *pb.CourseRequest, resp *pb.CourseListResponse) (err error) {
	resp.Code = config.SUCCESS
	// 查找备忘录
	r, count, err := dao.NewCourseDao(ctx).ListCourseByUserId(req.Uid, uint32(req.Week))
	if err != nil {
		resp.Code = config.ERROR
		// log.LogrusObj.Error("ListCourseByUserId err:%v", err)
		return
	}
	// 返回proto里面定义的类型
	var courseRes []*pb.CourseModel
	for _, item := range r {
		courseRes = append(courseRes, BuildCourse(item))
	}
	resp.CourseList = courseRes
	resp.Count = uint32(count)
	return
}

func (t *CourseSrv) UpdateCourse(ctx context.Context, req *pb.CourseRequest, resp *pb.CourseDetailResponse) (err error) {
	resp.Code = config.SUCCESS
	courseData, err := dao.NewCourseDao(ctx).UpdateCourse(req)
	if err != nil {
		resp.Code = config.ERROR
		// log.LogrusObj.Error("UpdateCourse err:%v", err)
		return
	}
	resp.CourseDetail = BuildCourse(courseData)
	return
}
func (t *CourseSrv) DeleteCourse(ctx context.Context, req *pb.CourseRequest, resp *pb.CourseDetailResponse) (err error) {
	resp.Code = config.SUCCESS
	err = dao.NewCourseDao(ctx).DeleteCourseByIdAndUserId(req.Id, req.Uid)
	if err != nil {
		resp.Code = config.ERROR
		// log.LogrusObj.Error("DeleteCourse err:%v", err)
		return
	}
	return
}

func BuildCourse(item *model.Course) *pb.CourseModel {
	courseModel := pb.CourseModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		CourseName: item.CourseName,
		StartClass: uint32(item.StartClass),
		EndClass:   uint32(item.EndClass),
		StartWeek:  uint32(item.StartWeek),
		EndWeek:    uint32(item.EndWeek),
		DayOfWeek:  uint32(item.DayOfWeek),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
	return &courseModel
}
