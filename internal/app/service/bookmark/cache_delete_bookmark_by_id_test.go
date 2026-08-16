package bookmark_test

import (
	"context"
	"testing"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	mock_cache "github.com/HemlockPham7/bookmark-service/internal/app/repository/cache/mocks"
	"github.com/HemlockPham7/bookmark-service/internal/app/service/bookmark"
	mock_bookmark "github.com/HemlockPham7/bookmark-service/internal/app/service/bookmark/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBookmarkService_DeleteBookmarkByID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupService func(ctx context.Context) *mock_bookmark.Service
		setupCache   func(ctx context.Context) *mock_cache.DB

		expectedResult *model.Bookmark
		expectedError  error
	}{
		{
			name: "delete bookmark successfully",

			setupService: func(ctx context.Context) *mock_bookmark.Service {
				mockService := mock_bookmark.NewService(t)
				mockService.On("DeleteBookmarkByID", ctx, "d7c13097-67a7-4eae-a60e-0b9b533b7bd4", "d7c13097-67a7-4eae-a60e-0b9b533b7bd6").Return(nil)
				return mockService
			},

			setupCache: func(ctx context.Context) *mock_cache.DB {
				mockCache := mock_cache.NewDB(t)
				mockCache.On("DeleteCache", ctx, "get_bookmarks_d7c13097-67a7-4eae-a60e-0b9b533b7bd4").Return(nil)
				return mockCache
			},

			expectedError: nil,
		},
		{
			name: "fail to delete bookmark",

			setupService: func(ctx context.Context) *mock_bookmark.Service {
				mockService := mock_bookmark.NewService(t)
				mockService.On("DeleteBookmarkByID", ctx, "d7c13097-67a7-4eae-a60e-0b9b533b7bd4", "d7c13097-67a7-4eae-a60e-0b9b533b7bd6").Return(assert.AnError)
				return mockService
			},

			setupCache: func(ctx context.Context) *mock_cache.DB {
				mockCache := mock_cache.NewDB(t)
				mockCache.On("DeleteCache", ctx, "get_bookmarks_d7c13097-67a7-4eae-a60e-0b9b533b7bd4").Return(assert.AnError)
				return mockCache
			},

			expectedError: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			mockService := tc.setupService(ctx)
			mockCache := tc.setupCache(ctx)

			testService := bookmark.NewBookmarkServiceWithCache(mockService, mockCache)
			err := testService.DeleteBookmarkByID(ctx, "d7c13097-67a7-4eae-a60e-0b9b533b7bd4", "d7c13097-67a7-4eae-a60e-0b9b533b7bd6")

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
