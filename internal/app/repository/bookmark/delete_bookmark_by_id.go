package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (r *bookmarkRepository) DeleteBookmarkByID(ctx context.Context, userID, bookmarkID string) error {
	span := newrelic.FromContext(ctx).StartSegment("DeleteBookmarkByID_BookmarkRepository")
	defer span.End()

	result := r.db.WithContext(ctx).Model(&model.Bookmark{}).Where("id = ? AND user_id = ?", bookmarkID, userID).Delete(&model.Bookmark{})

	if result.Error != nil {
		return dbutils.CatchDBError(result.Error)
	}

	if result.RowsAffected == 0 {
		return dbutils.ErrRecordNotFound
	}
	return nil
}
