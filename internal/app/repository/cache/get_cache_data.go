package cache

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (r *redisDB) GetCacheData(ctx context.Context, cacheGroupKey, cacheKey string) ([]byte, error) {
	span := newrelic.FromContext(ctx).StartSegment("GetCacheData_CacheRepository")
	defer span.End()

	return r.c.HGet(ctx, cacheGroupKey, cacheKey).Bytes()
}
