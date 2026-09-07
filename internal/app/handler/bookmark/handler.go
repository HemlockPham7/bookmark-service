package bookmark

import (
	"github.com/HemlockPham7/bookmark-service/internal/app/service/bookmark"
	"github.com/HemlockPham7/bookmark-service/internal/app/service/queue"
	"github.com/gin-gonic/gin"
)

// Handler defines the HTTP handlers for bookmark-related operations.
type Handler interface {
	CreateBookmark(c *gin.Context)
	UpdateBookmarkByID(c *gin.Context)
	DeleteBookmarkByID(c *gin.Context)
	GetBookmarks(c *gin.Context)
	ImportBookmarks(c *gin.Context)
}

type bookmarkHandler struct {
	bookmarkService bookmark.Service
	messageQueue    queue.Service
}

// NewHandler creates a new bookmark HTTP handler.
//
// Parameters:
//   - bookmarkService: the service used to handle bookmark operations.
//   - messageQueue: the message queue service used to process asynchronous tasks.
//
// Returns:
//   - A configured bookmark HTTP handler.
func NewHandler(bookmarkService bookmark.Service, messageQueue queue.Service) Handler {
	return &bookmarkHandler{bookmarkService: bookmarkService, messageQueue: messageQueue}
}
