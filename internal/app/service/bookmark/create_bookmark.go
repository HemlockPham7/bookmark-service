package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *bookmarkService) CreateBookmark(ctx context.Context, description, url, userID string) (*model.Bookmark, error) {
	span := newrelic.FromContext(ctx).StartSegment("CreateBookmark_BookmarkService")
	defer span.End()

	// create code
	code, err := s.codeGenerator.GenerateCode(codeLength)
	if err != nil {
		return nil, err
	}

	// create bookmark model
	bookmark := &model.Bookmark{
		Description: description,
		URL:         url,
		Code:        code,
		UserID:      userID,
	}

	// call repo
	res, err := s.bookmarkRepository.CreateBookmark(ctx, bookmark)
	if err != nil {
		return nil, err
	}

	//return
	return res, nil
}
