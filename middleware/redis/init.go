package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"mini-tiktok/config"
	"time"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})
}

func SetString(key string, value string) error {
	ctx := context.Background()
	err := rdb.Set(ctx, key, value, time.Duration(config.RedisExpireTime)*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetString 返回 key 对应的 value, bool 为是否存在
func GetString(key string) (string, bool, error) {
	ctx := context.Background()
	value, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", false, nil
	}

	if err != nil {
		return "", false, err
	}
	return value, true, nil
}
