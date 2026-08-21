package link

import (
	"context"
	"errors"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
)

var ErrCodeNotFound = errors.New("code not found")

func (s *linkRepository) GetURL(ctx context.Context, code string) (string, error) {
	span := newrelic.FromContext(ctx).StartSegment("GetURL_LinkRepository")
	defer span.End()

	res, err := s.c.Get(ctx, code).Result()
	if errors.Is(err, redis.Nil) {
		return "", ErrCodeNotFound
	}
	return res, err
}
