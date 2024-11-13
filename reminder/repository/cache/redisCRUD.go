package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func RemoveTasks(ctx context.Context, user_id uint64, deadline int64)  error {
	time_now := time.Now().Unix()
    float_dt := float64(time_now)-0.01
    user_id_string := strconv.FormatUint(uint64(user_id), 10)
    _, err := RedisClient.ZRemRangeByScore(ctx, user_id_string, "-inf", fmt.Sprintf("%f",float_dt)).Result()
    if err!= nil {
        return err
    }
    return nil
}

func GetTasks(ctx context.Context, user_id uint64, deadline int64) ([]redis.Z, error) {
	time_now := time.Now().Unix()
    deadline_time := time_now + deadline
    float_dt := float64(deadline_time)
    user_id_string := strconv.FormatUint(uint64(user_id), 10)
    tasks, err := RedisClient.ZRangeByScoreWithScores(ctx, user_id_string, &redis.ZRangeBy{
        Min: "-inf",
        Max: fmt.Sprintf("%f",float_dt)}).Result()
    if err!= nil {
        return nil, err 
    }     
    return tasks, nil
}