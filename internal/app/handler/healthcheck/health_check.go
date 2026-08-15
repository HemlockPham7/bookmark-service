package healthcheck

import (
	"net/http"

	"github.com/HemlockPham7/common-libs/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// HealthCheck checks the health of the service
// @Summary check redis health
// @Description ping and pong with redis server
// @Tags health-check
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.HealthCheckResponse
// @Router /health-check [get]
func (h *healthcheckHandler) HealthCheck(c *gin.Context) {
	msg, err := h.service.HealthCheck(c)
	if err != nil {
		log.Error().Err(err).Msg("Health-check error")
		c.JSON(http.StatusInternalServerError, response.InstanceErrResponse)
		return
	}
	c.JSON(http.StatusOK, msg)
}
