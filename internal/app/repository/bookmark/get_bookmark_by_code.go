package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
)

func (r *bookmarkRepository) GetBookmarkByCode(ctx context.Context, code string) (*model.Bookmark, error) {
	returnedBookmark := &model.Bookmark{}
	err := r.db.WithContext(ctx).Where("code = ?", code).First(returnedBookmark).Error
	if err != nil {
		return nil, dbutils.CatchDBError(err)
	}
	return returnedBookmark, nil
}
