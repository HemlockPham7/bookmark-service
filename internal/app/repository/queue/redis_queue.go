package queue

import "github.com/redis/go-redis/v9"

type redisQueue struct {
	client    *redis.Client
	queueName string
}

// NewRedisQueue creates a new Redis-based message queue repository.
//
// Parameters:
//   - c: the Redis client used to communicate with Redis.
//   - queueName: the name of the Redis queue.
//
// Returns:
//   - A Redis-based message queue repository.
func NewRedisQueue(c *redis.Client, queueName string) Repository {
	return &redisQueue{
		client:    c,
		queueName: queueName,
	}
}
