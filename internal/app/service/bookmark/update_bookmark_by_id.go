package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *bookmarkService) UpdateBookmarkByID(ctx context.Context, description, url, uid, bookmarkID string) (*model.Bookmark, error) {
	span := newrelic.FromContext(ctx).StartSegment("UpdateBookmarkByID_BookmarkService")
	defer span.End()

	updatedBookmark := &model.Bookmark{
		Description: description,
		URL:         url,
	}
	return s.bookmarkRepository.UpdateBookmarkByID(ctx, updatedBookmark, uid, bookmarkID)
}
