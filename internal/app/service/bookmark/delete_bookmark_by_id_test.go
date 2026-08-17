package bookmark

import (
	"context"
	"testing"

	mock_bookmark "github.com/HemlockPham7/bookmark-service/internal/app/repository/bookmark/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBookmarkService_DeleteBookmarkByID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMockRepo func(ctx context.Context, inputUserID, inputBookmarkID string) *mock_bookmark.Repository

		inputUserID     string
		inputBookmarkID string

		expectedError error
	}{
		{
			name: "success",

			setupMockRepo: func(ctx context.Context, inputUserID, inputBookmarkID string) *mock_bookmark.Repository {
				repoMock := mock_bookmark.NewRepository(t)
				repoMock.On("DeleteBookmarkByID", ctx, inputUserID, inputBookmarkID).Return(nil)
				return repoMock
			},

			inputUserID:     "d7c13097-67a7-4eae-a60e-0b9b533b7bd4",
			inputBookmarkID: "d7c13097-67a7-4eae-a60e-0b9b533b7bd6",

			expectedError: nil,
		},
		{
			name: "delete bookmarks failed - repository error",

			setupMockRepo: func(ctx context.Context, inputUserID, inputBookmarkID string) *mock_bookmark.Repository {
				repoMock := mock_bookmark.NewRepository(t)
				repoMock.On("DeleteBookmarkByID", ctx, inputUserID, inputBookmarkID).Return(assert.AnError)
				return repoMock
			},

			inputUserID:     "d7c13097-67a7-4eae-a60e-0b9b533b7bd4",
			inputBookmarkID: "d7c13097-67a7-4eae-a60e-0b9b533b7bd6",

			expectedError: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			repoMock := tc.setupMockRepo(ctx, tc.inputUserID, tc.inputBookmarkID)
			service := NewService(repoMock, nil)
			err := service.DeleteBookmarkByID(ctx, tc.inputUserID, tc.inputBookmarkID)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
