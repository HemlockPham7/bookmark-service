package link

import (
	"context"
	"errors"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
)

var ErrCodeNotFound = errors.New("code not found")

// GetURL retrieves the original URL associated with the specified shortened link code.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - code: the shortened link code used to retrieve the URL.
//
// Returns:
//   - The original URL associated with the code.
//   - ErrCodeNotFound if the code does not exist.
//   - An error if the URL cannot be retrieved from Redis.
func (s *linkRepository) GetURL(ctx context.Context, code string) (string, error) {
	span := newrelic.FromContext(ctx).StartSegment("GetURL_LinkRepository")
	defer span.End()

	res, err := s.c.Get(ctx, code).Result()
	if errors.Is(err, redis.Nil) {
		return "", ErrCodeNotFound
	}
	return res, err
}
