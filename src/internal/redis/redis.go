package redis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client interface {
	Get(context.Context, string) ([]byte, bool, error)
	Set(context.Context, string, []byte, time.Duration) error
}

type redisClient struct {
	client *redis.Client
}

func Initialize() Client {
	// Initialize Redis client
	client := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_ADDR"),
		Password:     os.Getenv("REDIS_PASSWORD"),
		DB:           0,
	})

	return &redisClient{
		client: client,
	}
}

func (c *redisClient) Get(ctx context.Context, key string) ([]byte, bool, error) {
	result, err := c.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, fmt.Errorf("failed to get from Redis: %w", err)
	}
	return result, true, nil
}

func (c *redisClient) Set(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	if err := c.client.Set(ctx, key, value, expiration).Err(); err != nil {
		return fmt.Errorf("failed to set to Redis: %w", err)
	}
	return nil
}
