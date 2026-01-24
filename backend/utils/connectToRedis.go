package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectToRedis() (*redis.Client, error) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_CONNECT_URI"))
	if err != nil {
		return nil, fmt.Errorf("Error Parsing URL! Error:%s", err)
	}

	RedisClient = redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = RedisClient.Set(ctx, "Ini", "New", 0).Err()
	if err != nil {
		return nil, fmt.Errorf("Error Setting Value! Error:%s", err)
	}

	_, err = RedisClient.Get(ctx, "Ini").Result()
	if err != nil {
		return nil, fmt.Errorf("Error Getting Value! Error:%s", err)
	}

	log.Infof("Redis Connection Successful: %s", os.Getenv("REDIS_CONNECT_URI"))
	return RedisClient, nil
}
