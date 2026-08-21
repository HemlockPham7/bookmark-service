package cache

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (r *redisDB) DeleteCache(ctx context.Context, key string) error {
	span := newrelic.FromContext(ctx).StartSegment("DeleteCache_CacheRepository")
	defer span.End()

	return r.c.Del(ctx, key).Err()
}
