package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// CreateBookmark creates a new bookmark in the database.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the database operation.
//   - bookmark: the bookmark to create.
//
// Returns:
//   - The created bookmark.
//   - An error if the bookmark cannot be created.
func (r *bookmarkRepository) CreateBookmark(ctx context.Context, bookmark *model.Bookmark) (*model.Bookmark, error) {
	span := newrelic.FromContext(ctx).StartSegment("CreateBookmark_BookmarkRepository")
	defer span.End()

	err := r.db.WithContext(ctx).Create(bookmark).Error
	if err != nil {
		return nil, dbutils.CatchDBError(err)
	}
	return bookmark, nil
}
