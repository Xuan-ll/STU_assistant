package taskconfig

import (
	"fmt"

	"gopkg.in/ini.v1"
)

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	Charset    string

	EtcdHost string
	EtcdPort string

	TaskServiceAddress string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDbName   int
)

func Init() {
	file, err := ini.Load("./task/taskconfig/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadMysqlData(file)
	LoadEtcd(file)
	LoadServer(file)
	LoadRedisData(file)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadServer(file *ini.File) {
	TaskServiceAddress = file.Section("server").Key("TaskServiceAddress").String()
}

func LoadRedisData(file *ini.File) {
	RedisHost = file.Section("redis").Key("RedisHost").String()
	RedisPort = file.Section("redis").Key("RedisPort").String()
	RedisPassword = file.Section("redis").Key("RedisPassword").String()
}
