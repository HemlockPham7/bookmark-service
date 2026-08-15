package link

import (
	"github.com/HemlockPham7/bookmark-service/internal/app/service/link"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	ShortenLink(c *gin.Context)
	Redirect(c *gin.Context)
}

type linkHandler struct {
	linkService link.Service
}

func NewLinkHandler(linkService link.Service) Handler {
	return &linkHandler{linkService: linkService}
}
