package bookmark

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *bookmarkService) DeleteBookmarkByID(ctx context.Context, userID, bookmarkID string) error {
	span := newrelic.FromContext(ctx).StartSegment("DeleteBookmarkByID_BookmarkService")
	defer span.End()

	return s.bookmarkRepository.DeleteBookmarkByID(ctx, userID, bookmarkID)
}
