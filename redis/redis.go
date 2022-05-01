package redis

import (
	"github.com/go-redis/redis"
	"time"
)

type Client struct {
	redisClient *redis.Client
	option      *Option
}

type Option struct {
	Addr     string
	Password string
	DB       int
}

func New(op *Option) *Client {
	redisOptions := &redis.Options{
		Addr:     op.Addr,
		Password: op.Password,
		DB:       op.DB,
	}
	return &Client{
		redisClient: redis.NewClient(redisOptions),
		option:      op,
	}
}

func (c *Client) GetKeys() ([]string, error) {
	cmd := c.redisClient.Keys("*")
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return cmd.Val(), nil
}

func (c *Client) GetKeyByPrefix(prefix, key string) (string, error) {
	cmd := c.redisClient.Get(key)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return cmd.Val(), nil
}

func (c *Client) SetKeyByPrefix(prefix, key string, value interface{}, expiration time.Duration) error {
	return c.redisClient.Set(prefix+key, value, expiration).Err()
}
