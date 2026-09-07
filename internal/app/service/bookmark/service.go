package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/bookmark-service/internal/app/repository/bookmark"
	"github.com/HemlockPham7/common-libs/pkg/utils"
)

const codeLength = 8

// Service defines the business logic for bookmark-related operations.
//
//go:generate mockery --name Service --filename service.go --outpkg mock_bookmark
type Service interface {
	CreateBookmark(ctx context.Context, description, url, userID string) (*model.Bookmark, error)
	UpdateBookmarkByID(ctx context.Context, description, url, userID, bookmarkID string) (*model.Bookmark, error)
	DeleteBookmarkByID(ctx context.Context, userID, ID string) error
	GetBookmarks(ctx context.Context, userID string, page, limit int) (*GetBookmarksResult, error)
}

type bookmarkService struct {
	bookmarkRepository bookmark.Repository
	codeGenerator      utils.GenCode
}

// NewService creates a new bookmark service.
//
// Parameters:
//   - bookmarkRepository: the repository used to persist and retrieve bookmarks.
//   - codeGenerator: the code generator used to generate unique bookmark codes.
//
// Returns:
//   - A configured bookmark service.
func NewService(bookmarkRepository bookmark.Repository, codeGenerator utils.GenCode) Service {
	return &bookmarkService{
		bookmarkRepository: bookmarkRepository,
		codeGenerator:      codeGenerator,
	}
}
