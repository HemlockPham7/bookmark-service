package gencode

import (
	"github.com/HemlockPham7/common-libs/pkg/utils"
	"github.com/gin-gonic/gin"
)

const codeLength = 12

type Handler interface {
	GenerateCode(c *gin.Context)
}

type genCodeHandler struct {
	genCodeService utils.GenCode
}

func NewHandler(genCodeService utils.GenCode) Handler {
	return &genCodeHandler{genCodeService: genCodeService}
}
