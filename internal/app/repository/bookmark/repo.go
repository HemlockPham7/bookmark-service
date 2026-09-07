package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"gorm.io/gorm"
)

// Repository defines the data access operations for bookmarks.
//
//go:generate mockery --name Repository --filename repo.go --outpkg mock_bookmark
type Repository interface {
	CreateBookmark(ctx context.Context, bookmark *model.Bookmark) (*model.Bookmark, error)
	GetBookmarks(ctx context.Context, userID string, limit, offset int) ([]*model.Bookmark, int64, error)
	UpdateBookmarkByID(ctx context.Context, updatedBookmark *model.Bookmark, userID, bookmarkID string) (*model.Bookmark, error)
	DeleteBookmarkByID(ctx context.Context, userID, bookmarkID string) error
	GetBookmarkByCode(ctx context.Context, code string) (*model.Bookmark, error)
}

type bookmarkRepository struct {
	db *gorm.DB
}

// NewRepository creates a new bookmark repository.
//
// Parameters:
//   - db: the GORM database client used to access bookmark data.
//
// Returns:
//   - A configured bookmark repository.
func NewRepository(db *gorm.DB) Repository {
	return &bookmarkRepository{db: db}
}
