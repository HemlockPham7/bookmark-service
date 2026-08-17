package bookmark

import (
	"testing"

	"github.com/HemlockPham7/bookmark-service/internal/integration_test/data/fixtures"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBookmarkRepository_DeleteBookmarkByID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func(t *testing.T) *gorm.DB

		inputUserID     string
		inputBookmarkID string

		expectedError error
	}{
		{
			name: "delete bookmark successfully",

			setupMock: func(t *testing.T) *gorm.DB {
				return fixtures.NewFixture(t, &fixtures.BookmarkCommonTestDB{})
			},

			inputUserID:     "d7c13097-67a7-4eae-a60e-0b9b533b7bd4",
			inputBookmarkID: "d7c13097-67a7-4eae-a60e-0b9b533b7bd6",
			expectedError:   nil,
		},
		{
			name: "failed to delete one bookmark with input code",

			setupMock: func(t *testing.T) *gorm.DB {
				return fixtures.NewFixture(t, &fixtures.BookmarkCommonTestDB{})
			},

			inputUserID:     "bookmark1",
			inputBookmarkID: "bookmark1",
			expectedError:   dbutils.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			dbMock := tc.setupMock(t)
			repo := NewRepository(dbMock)

			err := repo.DeleteBookmarkByID(ctx, tc.inputUserID, tc.inputBookmarkID)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
