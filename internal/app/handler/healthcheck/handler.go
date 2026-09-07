package healthcheck

import (
	"github.com/HemlockPham7/bookmark-service/internal/app/service/healthcheck"
	"github.com/gin-gonic/gin"
)

// Handler defines the HTTP handler for health check operations.
type Handler interface {
	HealthCheck(c *gin.Context)
}

type healthcheckHandler struct {
	service healthcheck.Service
}

// NewHealthcheckHandler creates a new health check HTTP handler.
//
// Parameters:
//   - service: the service used to perform health checks.
//
// Returns:
//   - A configured health check HTTP handler.
func NewHealthcheckHandler(service healthcheck.Service) Handler {
	return &healthcheckHandler{service: service}
}
