package bookmark

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/bookmark-service/internal/app/repository/bookmark"
	"github.com/HemlockPham7/common-libs/pkg/utils"
)

const codeLength = 8

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

func NewService(bookmarkRepository bookmark.Repository, codeGenerator utils.GenCode) Service {
	return &bookmarkService{
		bookmarkRepository: bookmarkRepository,
		codeGenerator:      codeGenerator,
	}
}
