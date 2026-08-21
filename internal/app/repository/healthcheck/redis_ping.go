package healthcheck

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (r *healthCheckRepository) RedisPing(ctx context.Context) error {
	span := newrelic.FromContext(ctx).StartSegment("RedisPing_HealthCheckRepository")
	defer span.End()

	return r.redisClient.Ping(ctx).Err()
}
