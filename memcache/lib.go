package memcache

import "time"

func Set(key string, value any, ttl time.Duration) {
	cache.Set(key, value, ttl)
}

func Get(key string) (any, bool) {
	return cache.Get(key)
}

func Remove(key string) {
	cache.Remove(key)
}
