package redis

import "github.com/go-redis/redis/v9"

type RedisClient struct {
	*redis.Client
}

func NewRedisClient(r *redis.Client) *RedisClient {
	return &RedisClient{r}
}
