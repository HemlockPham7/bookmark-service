package bookmark

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

// DeleteBookmarkByID deletes a bookmark by its ID and invalidates the user's bookmark cache.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - userID: the ID of the user who owns the bookmark.
//   - ID: the ID of the bookmark to delete.
//
// Returns:
//   - An error if the bookmark cannot be deleted by the underlying bookmark service.
func (s *bookmarkServiceWithCache) DeleteBookmarkByID(ctx context.Context, userID, ID string) error {
	span := newrelic.FromContext(ctx).StartSegment("DeleteBookmarkByID_BookmarkServiceWithCache")
	defer span.End()

	cacheGroupKey := fmt.Sprintf(getBookmarksCacheGroupKeyFormat, userID)
	err := s.c.DeleteCache(ctx, cacheGroupKey)
	if err != nil {
		log.Err(err).Str("key", cacheGroupKey).Msg("Failed to delete cache")
	}

	return s.s.DeleteBookmarkByID(ctx, userID, ID)
}
