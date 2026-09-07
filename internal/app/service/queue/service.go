package queue

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/repository/queue"
)

// Service defines the operations for processing queue-related jobs.
//
//go:generate mockery --name Service --filename service.go
type Service interface {
	SendImportBookmarkJob(ctx context.Context, uid string, bookmarkInputs []*ImportBookmarkInput) error
}

type service struct {
	messageQueue queue.Repository
}

// NewService creates a new queue service.
//
// Parameters:
//   - messageQueue: the repository used to push messages to the queue.
//
// Returns:
//   - A configured queue service.
func NewService(messageQueue queue.Repository) Service {
	return &service{
		messageQueue: messageQueue,
	}
}
