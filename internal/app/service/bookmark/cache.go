package bookmark

import (
	"time"

	"github.com/HemlockPham7/bookmark-service/internal/app/repository/cache"
)

const (
	getBookmarksCacheGroupKeyFormat = "get_bookmarks_%s"
	getBookmarksCacheKeyFormat      = "%d_%d"
	getBookmarksCacheExp            = 24 * time.Hour
)

type bookmarkServiceWithCache struct {
	s Service
	c cache.DB
}

func NewBookmarkServiceWithCache(s Service, c cache.DB) Service {
	return &bookmarkServiceWithCache{s: s, c: c}
}
