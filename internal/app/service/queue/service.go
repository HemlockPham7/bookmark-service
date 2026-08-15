package queue

import (
	"context"

	"github.com/HemlockPham7/bookmark-service/internal/app/repository/queue"
)

//go:generate mockery --name Service --filename service.go
type Service interface {
	SendImportBookmarkJob(ctx context.Context, uid string, bookmarkInputs []*ImportBookmarkInput) error
}

type service struct {
	messageQueue queue.Repository
}

func NewService(messageQueue queue.Repository) Service {
	return &service{
		messageQueue: messageQueue,
	}
}
