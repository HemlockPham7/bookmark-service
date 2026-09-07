package bookmark

import (
	"context"
	"fmt"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

// UpdateBookmarkByID updates a bookmark and invalidates the user's bookmark cache.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - description: the updated description of the bookmark.
//   - url: the updated URL of the bookmark.
//   - userID: the ID of the user who owns the bookmark.
//   - bookmarkID: the ID of the bookmark to update.
//
// Returns:
//   - The updated bookmark.
//   - An error if the bookmark cannot be updated.
func (s *bookmarkServiceWithCache) UpdateBookmarkByID(ctx context.Context, description, url, userID, ID string) (*model.Bookmark, error) {
	span := newrelic.FromContext(ctx).StartSegment("UpdateBookmarkByID_BookmarkServiceWithCache")
	defer span.End()

	cacheGroupKey := fmt.Sprintf(getBookmarksCacheGroupKeyFormat, userID)
	err := s.c.DeleteCache(ctx, cacheGroupKey)
	if err != nil {
		log.Err(err).Str("key", cacheGroupKey).Msg("Failed to delete cache")
	}

	return s.s.UpdateBookmarkByID(ctx, description, url, userID, ID)
}
