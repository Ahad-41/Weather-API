package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	Client *redis.Client
}

func ConnectRedis(addr string, password string) *RedisDB {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Warning: Failed to connect to Redis at %s: %v\n", addr, err)
		fmt.Println("The application will continue but caching will be disabled")
		return nil
	}

	fmt.Println("Connected to Redis successfully at", addr)
	return &RedisDB{Client: client}
}
