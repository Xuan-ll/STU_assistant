package cache

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func AddOrUpdateRedis(ctx context.Context, userId uint, title string, context string, endTime int64) error {
	// 将title和context拼接成一个字符串
	task := title + "|" + context
	// 将user_id转为string类型T
	userIdString := strconv.FormatUint(uint64(userId), 10)
	err := RedisClient.ZAdd(ctx, userIdString, redis.Z{
		Score:  float64(endTime),
		Member: task,
	}).Err()
	return err
}

func RemoveRedis(ctx context.Context, userId uint, title string, context string) error {
	// 将title和context拼接成一个字符串
	task := title + "|" + context
	// 将user_id转为string类型T
	userIdString := strconv.FormatUint(uint64(userId), 10)
	err := RedisClient.ZRem(ctx, userIdString, task).Err()
	if err==redis.Nil {
		err = nil
	}
	return err
}
