package redis

import "github.com/redis/go-redis/v9"

type Client struct {
	*redis.Client
}

func New() (Client, func()) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return Client{redisClient}, func() { _ = redisClient.Close() }
}
