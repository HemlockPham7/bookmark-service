package healthcheck

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/bookmark-service/internal/app/repository/healthcheck"
)

//go:generate mockery --name Service --filename service.go --outpkg mockHealthCheck
type Service interface {
	HealthCheck(ctx context.Context) (*model.HealthCheckResponse, error)
}

type healthcheckService struct {
	serviceName           string
	instanceID            string
	healthCheckRepository healthcheck.Repository
}

func NewHealthCheckService(serviceName, instanceId string, healthCheckRepository healthcheck.Repository) Service {
	return &healthcheckService{
		serviceName:           serviceName,
		instanceID:            instanceId,
		healthCheckRepository: healthCheckRepository,
	}
}
