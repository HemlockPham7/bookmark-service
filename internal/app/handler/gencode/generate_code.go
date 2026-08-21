package gencode

import (
	"net/http"

	"github.com/HemlockPham7/common-libs/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

// GenerateCode godoc
// @Summary Generate a random code
// @Tags code
// @Produce json
// @Success 200 {string} string
// @Router /gencode [get]
func (g *genCodeHandler) GenerateCode(c *gin.Context) {
	span := newrelic.FromContext(c).StartSegment("GenerateCode_GenCodeHandler")
	defer span.End()

	code, err := g.genCodeService.GenerateCode(codeLength)
	if err != nil {
		log.Error().Err(err).Str("from", "handler.genCodeService.GenerateCode").Msg("Cannot generate code")
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.InternalErrResponse)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}
