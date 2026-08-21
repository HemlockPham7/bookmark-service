package bookmark

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

func (s *bookmarkServiceWithCache) GetBookmarks(ctx context.Context, userID string, page, limit int) (*GetBookmarksResult, error) {
	nrTransaction := newrelic.FromContext(ctx)
	span := nrTransaction.StartSegment("GetBookmarks_BookmarkServiceWithCache")
	defer span.End()

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
			nrTransaction.Application().RecordCustomEvent("CacheHit", map[string]interface{}{
				"endpoint":  "GET /v1/bookmarks/get",
				"cache_hit": true,
			})
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
	nrTransaction.Application().RecordCustomEvent("CacheHit", map[string]interface{}{
		"endpoint":  "GET /v1/bookmarks/get",
		"cache_hit": false,
	})
	// return result
	return result, nil
}
