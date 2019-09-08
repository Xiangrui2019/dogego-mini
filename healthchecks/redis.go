package healthchecks

import "dogego-mini/cache"

func RedisHealthCheck() error {
	return cache.CacheClient.Ping().Err()
}
