package gencode

import (
	"github.com/HemlockPham7/common-libs/pkg/utils"
	"github.com/gin-gonic/gin"
)

const codeLength = 12

// Handler defines the HTTP handler for code generation operations.
type Handler interface {
	GenerateCode(c *gin.Context)
}

type genCodeHandler struct {
	genCodeService utils.GenCode
}

// NewHandler creates a new code generation HTTP handler.
//
// Parameters:
//   - genCodeService: the service used to generate codes.
//
// Returns:
//   - A configured code generation HTTP handler.
func NewHandler(genCodeService utils.GenCode) Handler {
	return &genCodeHandler{genCodeService: genCodeService}
}
