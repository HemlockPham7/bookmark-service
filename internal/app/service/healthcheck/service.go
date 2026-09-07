package healthcheck

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/bookmark-service/internal/app/repository/healthcheck"
)

// Service defines the business logic for service health checks.
//
//go:generate mockery --name Service --filename service.go --outpkg mockHealthCheck
type Service interface {
	HealthCheck(ctx context.Context) (*model.HealthCheckResponse, error)
}

type healthcheckService struct {
	serviceName           string
	instanceID            string
	healthCheckRepository healthcheck.Repository
}

// NewHealthCheckService creates a new health check service.
//
// Parameters:
//   - serviceName: the name of the service being monitored.
//   - instanceID: the unique identifier of the service instance.
//   - healthCheckRepository: the repository used to check service dependencies.
//
// Returns:
//   - A configured health check service.
func NewHealthCheckService(serviceName, instanceId string, healthCheckRepository healthcheck.Repository) Service {
	return &healthcheckService{
		serviceName:           serviceName,
		instanceID:            instanceId,
		healthCheckRepository: healthCheckRepository,
	}
}
