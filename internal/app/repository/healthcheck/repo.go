package healthcheck

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Repository defines the data access operations required for health checks.
//
//go:generate mockery --name Repository --filename repo.go --outpkg mocksHealthCheck
type Repository interface {
	RedisPing(ctx context.Context) error
}

type healthCheckRepository struct {
	redisClient *redis.Client
}

// NewHealthCheckRepository creates a new health check repository.
//
// Parameters:
//   - c: the Redis client used to check Redis connectivity.
//
// Returns:
//   - A configured health check repository.
func NewHealthCheckRepository(c *redis.Client) Repository {
	return &healthCheckRepository{redisClient: c}
}
