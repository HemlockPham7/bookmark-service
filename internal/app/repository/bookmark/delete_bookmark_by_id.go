package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
)

func (r *bookmarkRepository) DeleteBookmarkByID(ctx context.Context, userID, bookmarkID string) error {
	err := r.db.WithContext(ctx).Model(&model.Bookmark{}).Where("id = ? AND user_id = ?", bookmarkID, userID).Delete(&model.Bookmark{}).Error
	if err != nil {
		return dbutils.CatchDBError(err)
	}
	return nil
}
