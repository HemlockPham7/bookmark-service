package healthcheck

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
)

func (s *healthcheckService) HealthCheck(ctx context.Context) (*model.HealthCheckResponse, error) {
	if err := s.healthCheckRepository.RedisPing(ctx); err != nil {
		return nil, err
	}
	return &model.HealthCheckResponse{
		Message:     "OK",
		ServiceName: s.serviceName,
		InstanceID:  s.instanceID,
	}, nil
}
