package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mdportnov/common/util"
)

var ctx = context.Background()
var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr: util.GetEnv("REDIS_ADDR", "localhost:6379"),
	})
}

func SetCache(key string, value interface{}) error {
	return client.Set(ctx, key, value, 0).Err()
}

func GetCache(key string) (string, error) {
	return client.Get(ctx, key).Result()
}
