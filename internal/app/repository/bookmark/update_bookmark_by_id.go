package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm/clause"
)

// UpdateBookmarkByID updates a bookmark by its ID for the specified user.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the database operation.
//   - updatedBookmark: the bookmark fields to update.
//   - userID: the ID of the user who owns the bookmark.
//   - bookmarkID: the ID of the bookmark to update.
//
// Returns:
//   - The updated bookmark.
//   - An error if the bookmark cannot be updated or does not exist.
func (r *bookmarkRepository) UpdateBookmarkByID(ctx context.Context, updatedBookmark *model.Bookmark, userID, bookmarkID string) (*model.Bookmark, error) {
	span := newrelic.FromContext(ctx).StartSegment("UpdateBookmarkByID_BookmarkRepository")
	defer span.End()

	returnedBookmark := &model.Bookmark{}

	result := r.db.WithContext(ctx).
		Model(returnedBookmark).
		Clauses(clause.Returning{}).
		Where("id = ? AND user_id = ?", bookmarkID, userID).
		Updates(updatedBookmark)

	if result.Error != nil {
		return nil, dbutils.CatchDBError(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, dbutils.ErrRecordNotFound
	}

	return returnedBookmark, nil
}
