package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// CreateBookmark creates a new bookmark for the specified user.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - description: the description of the bookmark.
//   - url: the URL of the bookmark.
//   - userID: the ID of the user who owns the bookmark.
//
// Returns:
//   - The newly created bookmark.
//   - An error if the bookmark code cannot be generated or the bookmark cannot be created.
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
