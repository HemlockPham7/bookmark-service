package cache

import (
	"context"
	"testing"
	"time"

	redisPkg "github.com/HemlockPham7/common-libs/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisDB_GetCacheData(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func(ctx context.Context, t *testing.T) *redis.Client

		expectedError error

		verifyFunc func(ctx context.Context, redisClient *redis.Client, needTestedValue []byte)
	}{
		{
			name: "successful cache storage",

			setupMock: func(ctx context.Context, t *testing.T) *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)
				redisClient.HSet(ctx, "cache_group_key", "cache_key", []byte("cache_value"))
				redisClient.Expire(ctx, "cache_group_key", time.Hour)
				return redisClient
			},

			expectedError: nil,

			verifyFunc: func(ctx context.Context, redisClient *redis.Client, needTestedValue []byte) {
				cacheValue, err := redisClient.HGet(ctx, "cache_group_key", "cache_key").Bytes()
				assert.Nil(t, err)
				assert.Equal(t, needTestedValue, cacheValue)
			},
		},
		{
			name: "successful cache storage",

			setupMock: func(ctx context.Context, t *testing.T) *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)
				_ = redisClient.Close()
				return redisClient
			},

			expectedError: redis.ErrClosed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			redisClient := tc.setupMock(ctx, t)
			storage := NewRedisDB(redisClient)

			cacheValue, err := storage.GetCacheData(ctx, "cache_group_key", "cache_key")
			assert.Equal(t, tc.expectedError, err)

			if err == nil {
				tc.verifyFunc(ctx, redisClient, cacheValue)
			}

		})
	}
}
