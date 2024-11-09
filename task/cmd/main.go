package main

import (
	"fmt"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"stu_Assistant/idl/pb"
	"stu_Assistant/task/service"
	"stu_Assistant/task/repository/ormdb"
	"stu_Assistant/task/taskconfig"
)

func main() {
	taskconfig.Init()
	ormdb.InitDB()

    // etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", taskconfig.EtcdHost, taskconfig.EtcdPort)),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcTaskService"), // 微服务名字
		micro.Address(taskconfig.TaskServiceAddress),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterTaskServiceHandler(microService.Server(), service.GetTaskSrv())
	// 启动微服务
	_ = microService.Run()
}
	