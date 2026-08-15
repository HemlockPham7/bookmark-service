package cache

import "context"

func (r *redisDB) GetCacheData(ctx context.Context, cacheGroupKey, cacheKey string) ([]byte, error) {
	return r.c.HGet(ctx, cacheGroupKey, cacheKey).Bytes()
}
