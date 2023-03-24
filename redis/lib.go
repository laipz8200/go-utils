package cache

import (
	"fmt"

	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to redis: %v", err))
	}

	return &RedisCache{client: client}
}

var ErrEmptyKey = fmt.Errorf("key is empty")

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("failed to get key %s: %v", key, err)
	}

	return val, nil
}

func (c *RedisCache) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	if key == "" {
		return ErrEmptyKey
	}
	err := c.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key %s: %v", key, err)
	}

	return nil
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	if key == "" {
		return ErrEmptyKey
	}
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete key %s: %v", key, err)
	}

	return nil
}
