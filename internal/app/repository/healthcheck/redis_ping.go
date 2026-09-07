package healthcheck

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// RedisPing checks the connectivity to the Redis server.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//
// Returns:
//   - An error if the Redis server cannot be reached or does not respond successfully.
func (r *healthCheckRepository) RedisPing(ctx context.Context) error {
	span := newrelic.FromContext(ctx).StartSegment("RedisPing_HealthCheckRepository")
	defer span.End()

	return r.redisClient.Ping(ctx).Err()
}
