package bookmark

import (
	"context"
	"fmt"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

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
