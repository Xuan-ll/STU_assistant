package cache

import (
	"context"
	"fmt"
	"stu_Assistant/reminder/reminderconfig"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	host := reminderconfig.RedisHost
	port := reminderconfig.RedisPort
	password := reminderconfig.RedisPassword
    database := reminderconfig.RedisDbName
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