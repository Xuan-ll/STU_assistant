package main

import (
	"fmt"
	"time"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"

	"stu_Assistant/gateway/gwconfig"
	"stu_Assistant/gateway/router"
	"stu_Assistant/gateway/rpc"
)

func main() {
	gwconfig.Init()
	rpc.InitRPC()
	//cache.InitCache()
	//log.InitLog()
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", gwconfig.EtcdHost, gwconfig.EtcdPort)),
	)

	// 创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address(":4000"),
		// 将服务调用实例使用gin处理
		web.Handler(router.NewRouter()),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	// 接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
