package rediscache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ErrRecordNotFound = errors.New("cache record not found")

// redisCache represents a Redis cache client.
type redisCache struct {
	client *redis.Client
	ctx    context.Context
}

// WithContext returns a new Redis cache client with the given context.
func (c *redisCache) WithContext(ctx context.Context) *redisCache {
	return &redisCache{
		client: c.client,
		ctx:    ctx,
	}
}

// Set sets a key-value pair in the Redis cache.
func (c *redisCache) Set(key string, value interface{}, ttl time.Duration) error {
	return c.client.Set(c.ctx, key, value, ttl).Err()
}

// Get gets the value associated with the given key in the Redis cache.
func (c *redisCache) Get(key string) (string, error) {
	val, err := c.client.Get(c.ctx, key).Result()
	if err == redis.Nil {
		return "", ErrRecordNotFound
	}
	return val, err
}

// Remove deletes the given key from the Redis cache.
func (c *redisCache) Remove(key string) error {
	return c.client.Del(c.ctx, key).Err()
}

// NewCache creates a new Redis cache client.
func NewCache(addr, password string, db int) *redisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return &redisCache{
		client: client,
		ctx:    context.Background(),
	}
}

// NewCache creates a new Redis cache client with dsn.
func NewCacheWithDSN(dsn string) *redisCache {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return &redisCache{
		client: client,
		ctx:    context.Background(),
	}
}
