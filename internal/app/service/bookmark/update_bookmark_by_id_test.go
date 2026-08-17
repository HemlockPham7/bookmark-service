package bookmark

import (
	"context"
	"testing"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	mock_bookmark "github.com/HemlockPham7/bookmark-service/internal/app/repository/bookmark/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBookmarkService_UpdateBookmarkByID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMockRepo func(ctx context.Context, inputDescription, inputURL, inputUserID, inputBookmarkID string) *mock_bookmark.Repository

		inputDescription string
		inputURL         string
		inputUserID      string
		inputBookmarkID  string

		expectedError    error
		expectedBookmark *model.Bookmark
	}{
		{
			name: "success",

			setupMockRepo: func(ctx context.Context, inputDescription, inputURL, inputUserID, inputBookmarkID string) *mock_bookmark.Repository {
				repoMock := mock_bookmark.NewRepository(t)
				repoMock.On("UpdateBookmarkByID", ctx, &model.Bookmark{
					Description: inputDescription,
					URL:         inputURL,
				}, inputUserID, inputBookmarkID).Return(&model.Bookmark{
					Description: inputDescription,
					URL:         inputURL,
				}, nil)
				return repoMock
			},

			inputDescription: "Updated Description",
			inputURL:         "Updated URL",
			inputUserID:      "d7c13097-67a7-4eae-a60e-0b9b533b7bd4",
			inputBookmarkID:  "d7c13097-67a7-4eae-a60e-0b9b533b7bd6",

			expectedError: nil,
			expectedBookmark: &model.Bookmark{
				Description: "Updated Description",
				URL:         "Updated URL",
			},
		},
		{
			name: "get bookmarks failed - repository error",

			setupMockRepo: func(ctx context.Context, inputDescription, inputURL, inputUserID, inputBookmarkID string) *mock_bookmark.Repository {
				repoMock := mock_bookmark.NewRepository(t)
				repoMock.On("UpdateBookmarkByID", ctx, &model.Bookmark{
					Description: inputDescription,
					URL:         inputURL,
				}, inputUserID, inputBookmarkID).Return(nil, assert.AnError)
				return repoMock
			},

			inputDescription: "Updated Description",
			inputURL:         "Updated URL",
			inputUserID:      "d7c13097-67a7-4eae-a60e-0b9b533b7bd4",
			inputBookmarkID:  "d7c13097-67a7-4eae-a60e-0b9b533b7bd6",

			expectedError: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			repoMock := tc.setupMockRepo(ctx, tc.inputDescription, tc.inputURL, tc.inputUserID, tc.inputBookmarkID)
			service := NewService(repoMock, nil)
			updatedBookmarksResult, err := service.UpdateBookmarkByID(ctx, tc.inputDescription, tc.inputURL, tc.inputUserID, tc.inputBookmarkID)
			assert.Equal(t, tc.expectedError, err)
			if err == nil {
				assert.Equal(t, tc.expectedBookmark.Description, updatedBookmarksResult.Description)
				assert.Equal(t, tc.expectedBookmark.URL, updatedBookmarksResult.URL)
			}
		})
	}
}
