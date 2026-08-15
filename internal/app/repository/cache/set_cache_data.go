package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *redisDB) SetCacheData(ctx context.Context, cacheGroupKey, cacheKey string, value []byte, exp time.Duration) error {
	_, err := r.c.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		r.c.HSet(ctx, cacheGroupKey, cacheKey, value)
		r.c.Expire(ctx, cacheGroupKey, exp)
		return nil
	})

	return err
}
