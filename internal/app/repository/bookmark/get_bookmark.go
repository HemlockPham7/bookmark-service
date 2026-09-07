package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// GetBookmarks retrieves a paginated list of bookmarks and the total bookmark count for the specified user.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the database operations.
//   - userID: the ID of the user whose bookmarks are being retrieved.
//   - limit: the maximum number of bookmarks to return.
//   - offset: the number of bookmarks to skip before retrieving results.
//
// Returns:
//   - A list of bookmarks for the requested page.
//   - The total number of bookmarks owned by the user.
//   - An error if the bookmarks or total count cannot be retrieved.
func (r *bookmarkRepository) GetBookmarks(ctx context.Context, userID string, limit, offset int) ([]*model.Bookmark, int64, error) {
	span := newrelic.FromContext(ctx).StartSegment("GetBookmarks_BookmarkRepository")
	defer span.End()

	bookmarks, err := r.getBookmarks(ctx, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := r.countBookmarks(ctx, userID)
	if err != nil {
		return nil, 0, err
	}

	return bookmarks, total, nil
}

// getBookmarks retrieves a paginated list of bookmarks for the specified user.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the database operation.
//   - userID: the ID of the user whose bookmarks are being retrieved.
//   - limit: the maximum number of bookmarks to return.
//   - offset: the number of bookmarks to skip before retrieving results.
//
// Returns:
//   - A list of bookmarks ordered by creation time.
//   - An error if the bookmarks cannot be retrieved from the database.
func (r *bookmarkRepository) getBookmarks(ctx context.Context, userID string, limit, offset int) ([]*model.Bookmark, error) {
	span := newrelic.FromContext(ctx).StartSegment("GetBookmarksgetBookmarks_BookmarkRepository")
	defer span.End()

	bookmarks := make([]*model.Bookmark, 0, limit)

	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at ASC").Offset(offset).Limit(limit).Find(&bookmarks).Error
	if err != nil {
		return nil, dbutils.CatchDBError(err)
	}

	return bookmarks, nil
}

// countBookmarks counts the total number of bookmarks owned by the specified user.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the database operation.
//   - userID: the ID of the user whose bookmarks are being counted.
//
// Returns:
//   - The total number of bookmarks owned by the user.
//   - An error if the bookmark count cannot be retrieved from the database.
func (r *bookmarkRepository) countBookmarks(ctx context.Context, userID string) (int64, error) {
	span := newrelic.FromContext(ctx).StartSegment("GetBookmarkscountBookmarks_BookmarkRepository")
	defer span.End()

	var total int64

	err := r.db.WithContext(ctx).Model(&model.Bookmark{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return 0, dbutils.CatchDBError(err)
	}

	return total, nil
}
