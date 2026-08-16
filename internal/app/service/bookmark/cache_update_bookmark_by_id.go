package bookmark

import (
	"context"
	"fmt"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/rs/zerolog/log"
)

func (s *bookmarkServiceWithCache) UpdateBookmarkByID(ctx context.Context, description, url, userID, ID string) (*model.Bookmark, error) {
	cacheGroupKey := fmt.Sprintf(getBookmarksCacheGroupKeyFormat, userID)
	err := s.c.DeleteCache(ctx, cacheGroupKey)
	if err != nil {
		log.Err(err).Str("key", cacheGroupKey).Msg("Failed to delete cache")
	}

	return s.s.UpdateBookmarkByID(ctx, description, url, userID, ID)
}
