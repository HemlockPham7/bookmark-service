package bookmark

import (
	"github.com/HemlockPham7/bookmark-service/internal/app/service/bookmark"
	"github.com/HemlockPham7/bookmark-service/internal/app/service/queue"
	"github.com/gin-gonic/gin"
)

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

func NewHandler(bookmarkService bookmark.Service, messageQueue queue.Service) Handler {
	return &bookmarkHandler{bookmarkService: bookmarkService, messageQueue: messageQueue}
}
