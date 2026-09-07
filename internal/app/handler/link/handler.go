package link

import (
	"github.com/HemlockPham7/bookmark-service/internal/app/service/link"
	"github.com/gin-gonic/gin"
)

// Handler defines the HTTP handlers for link-related operations.
type Handler interface {
	ShortenLink(c *gin.Context)
	Redirect(c *gin.Context)
}

type linkHandler struct {
	linkService link.Service
}

// NewLinkHandler creates a new link HTTP handler.
//
// Parameters:
//   - linkService: the service used to handle link-related operations.
//
// Returns:
//   - A configured link HTTP handler.
func NewLinkHandler(linkService link.Service) Handler {
	return &linkHandler{linkService: linkService}
}
