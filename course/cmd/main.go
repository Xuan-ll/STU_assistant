package main

import (
	"fmt"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"stu_Assistant/course/repository/db/dao"
	// "stu_Assistant/course/repository/mq"

	"stu_Assistant/course/service"

	"stu_Assistant/course/config"
	"stu_Assistant/idl/pb"
	// log "github.com/CocaineCong/micro-todoList/pkg/logger"
)

func main() {
	config.Init()
	dao.InitDB()
	// mq.InitRabbitMQ()
	// log.InitLog()

	// 启动一些脚本
	// loadingScript()

	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcCourseService"), // 微服务名字
		micro.Address(config.CourseServiceAddress),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterCourseServiceHandler(microService.Server(), service.GetCourseSrv())
	// 启动微服务
	_ = microService.Run()
}

// func loadingScript() {
// 	ctx := context.Background()
// 	go script.CourseCreateSync(ctx)
// }
