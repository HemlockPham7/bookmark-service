package cache

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
)

// SetCacheData stores a value in a Redis hash and sets the expiration time for the cache group.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - cacheGroupKey: the Redis hash key used to group related cache entries.
//   - cacheKey: the key of the cache entry within the cache group.
//   - value: the value to store in the cache.
//   - exp: the expiration duration of the cache group.
//
// Returns:
//   - An error if the cache data cannot be stored or the expiration cannot be set.
func (r *redisDB) SetCacheData(ctx context.Context, cacheGroupKey, cacheKey string, value []byte, exp time.Duration) error {
	span := newrelic.FromContext(ctx).StartSegment("SetCacheData_CacheRepository")
	defer span.End()

	_, err := r.c.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		r.c.HSet(ctx, cacheGroupKey, cacheKey, value)
		r.c.Expire(ctx, cacheGroupKey, exp)
		return nil
	})

	return err
}
