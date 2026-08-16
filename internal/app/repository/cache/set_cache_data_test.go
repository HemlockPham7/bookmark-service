package cache

import (
	"context"
	"testing"
	"time"

	redisPkg "github.com/HemlockPham7/common-libs/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisDB_SetCacheData(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func(ctx context.Context, t *testing.T) *redis.Client

		expectedError error

		verifyFunc func(ctx context.Context, redisClient *redis.Client)
	}{
		{
			name: "successful cache storage",

			setupMock: func(ctx context.Context, t *testing.T) *redis.Client {
				return redisPkg.InitMockRedis(t)
			},

			expectedError: nil,

			verifyFunc: func(ctx context.Context, redisClient *redis.Client) {
				val, err := redisClient.HGet(ctx, "cache_group_key", "cache_key").Result()

				assert.Nil(t, err)

				assert.Equal(t, "cache_value", val)
			},
		},
		{
			name: "failed cache storage due to closed redis client",

			setupMock: func(ctx context.Context, t *testing.T) *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)
				_ = redisClient.Close()
				return redisClient
			},

			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			redisClient := tc.setupMock(ctx, t)
			storage := NewRedisDB(redisClient)

			err := storage.SetCacheData(ctx, "cache_group_key", "cache_key", []byte("cache_value"), time.Hour)
			assert.Equal(t, tc.expectedError, err)

			if tc.verifyFunc != nil {
				tc.verifyFunc(ctx, redisClient)
			}
		})
	}
}
