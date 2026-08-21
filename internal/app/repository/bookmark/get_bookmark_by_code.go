package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (r *bookmarkRepository) GetBookmarkByCode(ctx context.Context, code string) (*model.Bookmark, error) {
	span := newrelic.FromContext(ctx).StartSegment("GetBookmarkByCode_BookmarkRepository")
	defer span.End()

	returnedBookmark := &model.Bookmark{}
	err := r.db.WithContext(ctx).Where("code = ?", code).First(returnedBookmark).Error
	if err != nil {
		return nil, dbutils.CatchDBError(err)
	}
	return returnedBookmark, nil
}
