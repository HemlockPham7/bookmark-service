package bookmark

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// DeleteBookmarkByID deletes a bookmark by its ID for the specified user.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - userID: the ID of the user who owns the bookmark.
//   - bookmarkID: the ID of the bookmark to delete.
//
// Returns:
//   - An error if the bookmark cannot be deleted from the repository.
func (s *bookmarkService) DeleteBookmarkByID(ctx context.Context, userID, bookmarkID string) error {
	span := newrelic.FromContext(ctx).StartSegment("DeleteBookmarkByID_BookmarkService")
	defer span.End()

	return s.bookmarkRepository.DeleteBookmarkByID(ctx, userID, bookmarkID)
}
