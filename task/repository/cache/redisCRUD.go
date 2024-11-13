package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func AddOrUpdateRedis(ctx context.Context, userId uint, title string, startTime int64, endTime int64) error {
	t := time.Unix(startTime, 0)
	format := "2006-01-02 15:04:05"
	startTimeString := t.Format(format)
	// 将title和end time拼接成一个字符串
	task := title + "," + startTimeString
	// 将user_id转为string类型T
	userIdString := strconv.FormatUint(uint64(userId), 10)
	err := RedisClient.ZAdd(ctx, userIdString, redis.Z{
		Score:  float64(endTime),
		Member: task,
	}).Err()
	return err
}

func RemoveRedis(ctx context.Context, userId uint, title string, startTime int64, endTime int64) error {
	t := time.Unix(startTime, 0)
	format := "2006-01-02 15:04:05"
	startTimeString := t.Format(format)
	// 将title和end time拼接成一个字符串
	task := title + "," + startTimeString	
	// 将user_id转为string类型T
	userIdString := strconv.FormatUint(uint64(userId), 10)
	err := RedisClient.ZRem(ctx, userIdString, task).Err()
	if err==redis.Nil {
		err = nil
	}
	return err
}
