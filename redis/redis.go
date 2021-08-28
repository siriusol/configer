package redis

import (
	"github.com/go-redis/redis"
	"time"
)

var (
	redisClient *redis.Client
)

func init() {
	redisOptions := &redis.Options{
		Addr:               "127.0.0.1:6379",
		Password:           "",
		DB:                 0,
	}
	redisClient = redis.NewClient(redisOptions)
}

func GetKeys() ([]string, error) {
	cmd := redisClient.Keys("*")
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return cmd.Val(), nil
}

func GetKeyByPrefix(prefix, key string) (string, error) {
	cmd := redisClient.Get(key)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return cmd.Val(), nil
}

func SetKeyByPrefix(prefix, key string, value interface{}) error {
	return redisClient.Set(prefix + key, value, time.Second).Err()
}
