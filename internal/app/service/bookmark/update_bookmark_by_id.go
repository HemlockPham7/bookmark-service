package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// UpdateBookmarkByID updates the description and URL of a bookmark owned by the specified user.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - description: the updated description of the bookmark.
//   - url: the updated URL of the bookmark.
//   - userID: the ID of the user who owns the bookmark.
//   - bookmarkID: the ID of the bookmark to update.
//
// Returns:
//   - The updated bookmark.
//   - An error if the bookmark cannot be updated.
func (s *bookmarkService) UpdateBookmarkByID(ctx context.Context, description, url, uid, bookmarkID string) (*model.Bookmark, error) {
	span := newrelic.FromContext(ctx).StartSegment("UpdateBookmarkByID_BookmarkService")
	defer span.End()

	updatedBookmark := &model.Bookmark{
		Description: description,
		URL:         url,
	}
	return s.bookmarkRepository.UpdateBookmarkByID(ctx, updatedBookmark, uid, bookmarkID)
}
