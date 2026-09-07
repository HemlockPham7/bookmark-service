package bookmark

import (
	"context"
	"fmt"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

// CreateBookmark creates a new bookmark and invalidates the user's bookmark cache.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - description: the description of the bookmark.
//   - url: the URL of the bookmark.
//   - userID: the ID of the user who owns the bookmark.
//
// Returns:
//   - The newly created bookmark.
//   - An error if the underlying bookmark service fails to create the bookmark.
func (s *bookmarkServiceWithCache) CreateBookmark(ctx context.Context, description, url, userID string) (*model.Bookmark, error) {
	span := newrelic.FromContext(ctx).StartSegment("CreateBookmark_BookmarkServiceWithCache")
	defer span.End()

	cacheGroupKey := fmt.Sprintf(getBookmarksCacheGroupKeyFormat, userID)
	err := s.c.DeleteCache(ctx, cacheGroupKey)
	if err != nil {
		log.Err(err).Str("key", cacheGroupKey).Msg("Failed to delete cache")
	}

	return s.s.CreateBookmark(ctx, description, url, userID)
}
