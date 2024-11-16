package cache

import (
	"context"
	"fmt"
	"stu_Assistant/task/taskconfig"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	host := taskconfig.RedisHost
	port := taskconfig.RedisPort
	password := taskconfig.RedisPassword
    database := taskconfig.RedisDbName
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       database,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}