package rpc

import (
	"go-micro.dev/v4"

	"stu_Assistant/idl/pb"
)

var (
	UserService     pb.UserService
	TaskService     pb.TaskService
	CourseService   pb.CourseService
	ReminderService pb.ReminderService
)

func InitRPC() {
	// 用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
	)
	// 用户服务调用实例
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	// task
	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
	)
	taskService := pb.NewTaskService("rpcTaskService", taskMicroService.Client())

	// reminder
	reminderMicroService := micro.NewService(
		micro.Name("reminderService.client"),
	)
	reminderService := pb.NewReminderService("rpcReminderService", reminderMicroService.Client())

	CourseMicroService := micro.NewService(
		micro.Name("courseService.client"),
	)
	courseService := pb.NewCourseService("rpcCourseService", CourseMicroService.Client())

	UserService = userService
	TaskService = taskService
	ReminderService = reminderService
	CourseService = courseService
}
