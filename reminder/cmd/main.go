package main

import (
	"fmt"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"stu_Assistant/idl/pb"
	"stu_Assistant/reminder/reminderconfig"
	"stu_Assistant/reminder/repository/cache"
	"stu_Assistant/reminder/repository/ormdb"
	"stu_Assistant/reminder/service"
)

func main() {
	reminderconfig.Init()
	ormdb.InitDB()
    cache.InitRedis()
    // etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", reminderconfig.EtcdHost, reminderconfig.EtcdPort)),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcReminderService"), // 微服务名字
		micro.Address(reminderconfig.ReminderServiceAddress),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterReminderServiceHandler(microService.Server(), service.GetReminderSrv())
	// 启动微服务
	_ = microService.Run()
}