package cache

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// GetCacheData retrieves a cached value from a Redis hash.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - cacheGroupKey: the Redis hash key containing the cache entry.
//   - cacheKey: the key of the cache entry within the cache group.
//
// Returns:
//   - The cached value.
//   - An error if the cache entry cannot be retrieved or does not exist.
func (r *redisDB) GetCacheData(ctx context.Context, cacheGroupKey, cacheKey string) ([]byte, error) {
	span := newrelic.FromContext(ctx).StartSegment("GetCacheData_CacheRepository")
	defer span.End()

	return r.c.HGet(ctx, cacheGroupKey, cacheKey).Bytes()
}
