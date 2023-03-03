package rediscache

import (
	"context"
	"time"
)

var std *redisCache

// Init std object.
func Init(addr, password string, db int) {
	std = NewCache(addr, password, db)
}

// Init std object with dsn.
func InitWithDSN(dsn string) {
	std = NewCacheWithDSN(dsn)
}

// WithContext returns a new Redis cache client with the given context.
func WithContext(ctx context.Context) *redisCache {
	return std.WithContext(ctx)
}

// Set sets a key-value pair in the Redis cache.
func Set(key string, value interface{}, ttl time.Duration) error {
	return std.Set(key, value, ttl)
}

// Get gets the value associated with the given key in the Redis cache.
func Get(key string) (string, error) {
	return std.Get(key)
}

// Remove deletes the given key from the Redis cache.
func Remove(key string) error {
	return std.Remove(key)
}
