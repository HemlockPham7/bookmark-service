package link

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Repository defines the data access operations for shortened links.
//
//go:generate mockery --name Repository --filename repo.go --outpkg mockLink
type Repository interface {
	StoreURL(ctx context.Context, code, url string, expSecond int64) error
	GetURL(ctx context.Context, code string) (string, error)
}

type linkRepository struct {
	c *redis.Client
}

// NewLinkRepository creates a new Redis-based link repository.
//
// Parameters:
//   - c: the Redis client used to store and retrieve shortened links.
//
// Returns:
//   - A Redis-based link repository.
func NewLinkRepository(c *redis.Client) Repository {
	return &linkRepository{c: c}
}
