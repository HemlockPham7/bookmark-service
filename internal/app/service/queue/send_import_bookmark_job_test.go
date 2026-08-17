package queue

import (
	"context"
	"encoding/json"
	"testing"

	mock_queue "github.com/HemlockPham7/bookmark-service/internal/app/repository/queue/mocks"
	"github.com/stretchr/testify/assert"
)

func TestQueue_SendImportBookmarkJob(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		uid       string
		bookmarks []*ImportBookmarkInput

		setupMockRepo func(ctx context.Context, uid string, bookmarks []*ImportBookmarkInput) *mock_queue.Repository

		expectedError error
	}{
		{
			name: "successful send import bookmark job",

			uid: "user-123",
			bookmarks: []*ImportBookmarkInput{
				{
					URL:         "https://www.example.com",
					Description: "An example bookmark",
				},
			},

			setupMockRepo: func(ctx context.Context, uid string, bookmarks []*ImportBookmarkInput) *mock_queue.Repository {
				mockRepo := mock_queue.NewRepository(t)
				input := ImportMessage{
					UID:       uid,
					Bookmarks: bookmarks,
				}
				inputBytes, err := json.Marshal(input)
				assert.NoError(t, err)
				mockRepo.On("PushMessage", ctx, inputBytes).Return(nil)
				return mockRepo
			},

			expectedError: nil,
		},
		{
			name: "failed to send import bookmark job",

			uid: "user-123",
			bookmarks: []*ImportBookmarkInput{
				{
					URL:         "https://www.example.com",
					Description: "An example bookmark",
				},
			},

			setupMockRepo: func(ctx context.Context, uid string, bookmarks []*ImportBookmarkInput) *mock_queue.Repository {
				mockRepo := mock_queue.NewRepository(t)
				input := ImportMessage{
					UID:       uid,
					Bookmarks: bookmarks,
				}
				inputBytes, err := json.Marshal(input)
				assert.NoError(t, err)
				mockRepo.On("PushMessage", ctx, inputBytes).Return(assert.AnError)
				return mockRepo
			},

			expectedError: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			mockRepo := tc.setupMockRepo(ctx, tc.uid, tc.bookmarks)
			mockService := NewService(mockRepo)

			err := mockService.SendImportBookmarkJob(ctx, tc.uid, tc.bookmarks)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
