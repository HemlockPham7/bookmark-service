package bookmark

import (
	"errors"
	"testing"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/bookmark-service/internal/integration_test/data/fixtures"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBookmarkRepository_UpdateBookmarkByID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func(t *testing.T) *gorm.DB

		updatedBookmark *model.Bookmark
		inputUserID     string
		inputBookmarkID string

		expectedError    error
		expectedBookmark *model.Bookmark
	}{
		{
			name: "update one bookmark successfully",

			setupMock: func(t *testing.T) *gorm.DB {
				return fixtures.NewFixture(t, &fixtures.BookmarkCommonTestDB{})
			},

			updatedBookmark: &model.Bookmark{
				Description: "Updated Bookmark",
				URL:         "Updated URL",
			},
			inputUserID:     "d7c13097-67a7-4eae-a60e-0b9b533b7bd4",
			inputBookmarkID: "d7c13097-67a7-4eae-a60e-0b9b533b7bd6",
			expectedError:   nil,
			expectedBookmark: &model.Bookmark{
				Description: "Updated Bookmark",
				URL:         "Updated URL",
			},
		},
		{
			name: "invalid value, should be pointer to struct or slice",

			setupMock: func(t *testing.T) *gorm.DB {
				return fixtures.NewFixture(t, &fixtures.BookmarkCommonTestDB{})
			},

			inputUserID:     "bookmark1",
			inputBookmarkID: "bookmark1",
			expectedError:   errors.New("invalid value, should be pointer to struct or slice"),
		},
		{
			name: "failed to update one bookmark with invalid user id",

			setupMock: func(t *testing.T) *gorm.DB {
				return fixtures.NewFixture(t, &fixtures.BookmarkCommonTestDB{})
			},

			updatedBookmark: &model.Bookmark{
				Description: "Updated Bookmark",
				URL:         "Updated URL",
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

			bookmark, err := repo.UpdateBookmarkByID(ctx, tc.updatedBookmark, tc.inputUserID, tc.inputBookmarkID)

			assert.Equal(t, tc.expectedError, err)
			if err == nil {
				assert.Equal(t, tc.expectedBookmark.Description, bookmark.Description)
				assert.Equal(t, tc.expectedBookmark.URL, bookmark.URL)
			}
		})
	}
}
