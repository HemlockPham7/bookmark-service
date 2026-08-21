package bookmark

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

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
