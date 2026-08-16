package bookmark_test

import (
	"context"
	"testing"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	mock_cache "github.com/HemlockPham7/bookmark-service/internal/app/repository/cache/mocks"
	"github.com/HemlockPham7/bookmark-service/internal/app/service/bookmark"
	mock_bookmark "github.com/HemlockPham7/bookmark-service/internal/app/service/bookmark/mocks"
	"github.com/HemlockPham7/bookmark-service/internal/integration_test/data/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestBookmarkServiceWithCache_CreateBookmark(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupService func(ctx context.Context) *mock_bookmark.Service
		setupCache   func(ctx context.Context) *mock_cache.DB

		expectedResult *model.Bookmark
		expectedError  error
	}{
		{
			name: "create bookmark successfully",

			setupService: func(ctx context.Context) *mock_bookmark.Service {
				mockService := mock_bookmark.NewService(t)
				mockService.On("CreateBookmark", ctx, "abc xyz", "https://www.youtube.com", "d7c13097-a60e-67a7-4eae-0b9b533b7bd4").Return(&model.Bookmark{
					Base:        fixtures.GetTestBase("d7c13097-67a7-4eae-a60e-0b9b533b7bd6"),
					Description: "abc xyz",
					URL:         "https://www.youtube.com",
					Code:        "abcxyz",
					UserID:      "d7c13097-a60e-67a7-4eae-0b9b533b7bd4",
				}, nil)
				return mockService
			},

			setupCache: func(ctx context.Context) *mock_cache.DB {
				mockCache := mock_cache.NewDB(t)
				mockCache.On("DeleteCache", ctx, "get_bookmarks_d7c13097-a60e-67a7-4eae-0b9b533b7bd4").Return(nil)
				return mockCache
			},

			expectedResult: &model.Bookmark{
				Base:        fixtures.GetTestBase("d7c13097-67a7-4eae-a60e-0b9b533b7bd6"),
				Description: "abc xyz",
				URL:         "https://www.youtube.com",
				Code:        "abcxyz",
				UserID:      "d7c13097-a60e-67a7-4eae-0b9b533b7bd4",
			},
			expectedError: nil,
		},
		{
			name:           "fail to create bookmark",
			expectedResult: nil,
			expectedError:  assert.AnError,

			setupService: func(ctx context.Context) *mock_bookmark.Service {
				mockService := mock_bookmark.NewService(t)
				mockService.On("CreateBookmark", ctx, "abc xyz", "https://www.youtube.com", "d7c13097-a60e-67a7-4eae-0b9b533b7bd4").Return(nil, assert.AnError)
				return mockService
			},

			setupCache: func(ctx context.Context) *mock_cache.DB {
				mockCache := mock_cache.NewDB(t)
				mockCache.On("DeleteCache", ctx, "get_bookmarks_d7c13097-a60e-67a7-4eae-0b9b533b7bd4").Return(assert.AnError)
				return mockCache
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			mockService := tc.setupService(ctx)
			mockCache := tc.setupCache(ctx)

			testService := bookmark.NewBookmarkServiceWithCache(mockService, mockCache)
			bookmarkResult, err := testService.CreateBookmark(ctx, "abc xyz", "https://www.youtube.com", "d7c13097-a60e-67a7-4eae-0b9b533b7bd4")

			assert.Equal(t, tc.expectedResult, bookmarkResult)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
