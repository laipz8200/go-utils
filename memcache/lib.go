package memcache

func Set(key string, value any) {
	cache.setWithTTL(key, value, 0)
}

func SetWithTTL(key string, value any, ttl int64) {
	cache.setWithTTL(key, value, ttl)
}

func Get(key string) (any, bool) {
	return cache.get(key)
}

func Remove(key string) {
	cache.remove(key)
}
