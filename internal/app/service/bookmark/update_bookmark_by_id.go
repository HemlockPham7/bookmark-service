package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
)

func (s *bookmarkService) UpdateBookmarkByID(ctx context.Context, description, url, uid, bookmarkID string) (*model.Bookmark, error) {
	updatedBookmark := &model.Bookmark{
		Description: description,
		URL:         url,
	}
	return s.bookmarkRepository.UpdateBookmarkByID(ctx, updatedBookmark, uid, bookmarkID)
}
