package redis

import (
	"fmt"
	"os"

	redis "github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
)

func InitRedis() {
	var (
		redisAddress = os.Getenv("REDIS_HOST")
		redisPort    = os.Getenv("REDIS_PORT")
	)

	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisAddress, redisPort),
		DB:   0,
	})
}

func GetRedis() *redis.Client {
	if redisClient == nil {
		InitRedis()
	}

	return redisClient
}
