package healthcheck

import (
	"github.com/HemlockPham7/bookmark-service/internal/app/service/healthcheck"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	HealthCheck(c *gin.Context)
}

type healthcheckHandler struct {
	service healthcheck.Service
}

func NewHealthcheckHandler(service healthcheck.Service) Handler {
	return &healthcheckHandler{service: service}
}
