package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
)

type GetBookmarksResult struct {
	Bookmarks []*model.Bookmark `json:"bookmarks"`
	Total     int64             `json:"total"`
}

func (s *bookmarkService) GetBookmarks(ctx context.Context, userID string, page, limit int) (*GetBookmarksResult, error) {
	offset := (page - 1) * limit

	bookmarks, total, err := s.bookmarkRepository.GetBookmarks(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return &GetBookmarksResult{
		Bookmarks: bookmarks,
		Total:     total,
	}, nil
}
