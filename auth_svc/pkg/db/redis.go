package db

import (
	"auth_svc/pkg/config"
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func InitRedis(c *config.Config) (*redis.Client, error) {
	opt, _ := redis.ParseURL(c.RedisAddress)
  client := redis.NewClient(opt)
	// Test the connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("Failed to connect to Redis:", err)
	}

	return client, nil
}
