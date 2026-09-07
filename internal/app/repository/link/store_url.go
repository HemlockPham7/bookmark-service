package link

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// StoreURL stores a URL in Redis using the specified code as the key.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - code: the shortened link code used as the Redis key.
//   - url: the original URL to store.
//   - expSecond: the number of seconds before the stored URL expires.
//
// Returns:
//   - An error if the URL cannot be stored in Redis.
func (s *linkRepository) StoreURL(ctx context.Context, code, url string, expSecond int64) error {
	span := newrelic.FromContext(ctx).StartSegment("StoreURL_LinkRepository")
	defer span.End()

	return s.c.Set(ctx, code, url, time.Duration(expSecond)*time.Second).Err()
}
