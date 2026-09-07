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

// NewBookmarkServiceWithCache creates a new bookmark service with cache support.
//
// Parameters:
//   - s: the underlying bookmark service used to perform bookmark operations.
//   - c: the cache database used to store and invalidate bookmark data.
//
// Returns:
//   - A bookmark service with cache support.
func NewBookmarkServiceWithCache(s Service, c cache.DB) Service {
	return &bookmarkServiceWithCache{s: s, c: c}
}
