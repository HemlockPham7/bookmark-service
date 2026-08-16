package bookmark

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (s *bookmarkServiceWithCache) GetBookmarks(ctx context.Context, userID string, page, limit int) (*GetBookmarksResult, error) {
	// tao cache key
	cacheGroupKey := fmt.Sprintf(getBookmarksCacheGroupKeyFormat, userID)
	cacheKey := fmt.Sprintf(getBookmarksCacheKeyFormat, page, limit)

	// get cache data
	cacheData, err := s.c.GetCacheData(ctx, cacheGroupKey, cacheKey)

	// if cache exits, return cache
	if err == nil && len(cacheData) > 0 {
		result := &GetBookmarksResult{}

		err := json.Unmarshal(cacheData, result)
		if err != nil {
			cacheErr := s.c.DeleteCache(ctx, cacheGroupKey)
			if cacheErr != nil {
				log.Err(cacheErr).Str("key", cacheGroupKey).Msg("Failed to delete cache")
			}
		} else {
			return result, nil
		}
	}

	// if not, call service
	result, err := s.s.GetBookmarks(ctx, userID, page, limit)
	if err != nil {
		return nil, err
	}

	// save cache
	resultBytes, err := json.Marshal(result)
	if err == nil {
		cacheErr := s.c.SetCacheData(ctx, cacheGroupKey, cacheKey, resultBytes, getBookmarksCacheExp)
		if cacheErr != nil {
			log.Err(cacheErr).Str("key", cacheGroupKey).Msg("failed to set cache")
		}
	}

	// return result
	return result, nil
}
