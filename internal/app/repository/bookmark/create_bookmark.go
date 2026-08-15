package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
)

func (r *bookmarkRepository) CreateBookmark(ctx context.Context, bookmark *model.Bookmark) (*model.Bookmark, error) {
	err := r.db.WithContext(ctx).Create(bookmark).Error
	if err != nil {
		return nil, dbutils.CatchDBError(err)
	}
	return bookmark, nil
}
