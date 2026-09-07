package cache

import "github.com/redis/go-redis/v9"

type redisDB struct {
	c *redis.Client
}

// NewRedisDB creates a new Redis-based cache database.
//
// Parameters:
//   - c: the Redis client used to store and retrieve cached data.
//
// Returns:
//   - A Redis-based cache database.
func NewRedisDB(c *redis.Client) DB {
	return &redisDB{c: c}
}
