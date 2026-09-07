package healthcheck

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// HealthCheck checks the health of the service by verifying its Redis dependency.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//
// Returns:
//   - A health check response containing the service name and instance ID.
//   - An error if the Redis dependency is unavailable.
func (s *healthcheckService) HealthCheck(ctx context.Context) (*model.HealthCheckResponse, error) {
	span := newrelic.FromContext(ctx).StartSegment("HealthCheck_HealthCheckService")
	defer span.End()

	if err := s.healthCheckRepository.RedisPing(ctx); err != nil {
		return nil, err
	}
	return &model.HealthCheckResponse{
		Message:     "OK",
		ServiceName: s.serviceName,
		InstanceID:  s.instanceID,
	}, nil
}
