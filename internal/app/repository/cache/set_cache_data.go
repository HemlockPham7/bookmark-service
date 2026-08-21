package cache

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
)

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
