package initializers

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
)

func ConnectRedis() {
	opt, err := redis.ParseURL(LoadEnvVariable("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	redisClient = redis.NewClient(opt)

	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis connected successfully", pong)
}

func CloseRedis() {
	if redisClient != nil {
		err := redisClient.Close()
		if err != nil {
			fmt.Println("Error closing Redis connection:", err)
		}
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}
