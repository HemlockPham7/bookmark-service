package healthcheck

import (
	"context"

	"github.com/redis/go-redis/v9"
)

//go:generate mockery --name Repository --filename repo.go --outpkg mocksHealthCheck
type Repository interface {
	RedisPing(ctx context.Context) error
}

type healthCheckRepository struct {
	redisClient *redis.Client
}

func NewHealthCheckRepository(c *redis.Client) Repository {
	return &healthCheckRepository{redisClient: c}
}
