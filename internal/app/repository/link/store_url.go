package link

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *linkRepository) StoreURL(ctx context.Context, code, url string, expSecond int64) error {
	span := newrelic.FromContext(ctx).StartSegment("StoreURL_LinkRepository")
	defer span.End()

	return s.c.Set(ctx, code, url, time.Duration(expSecond)*time.Second).Err()
}
