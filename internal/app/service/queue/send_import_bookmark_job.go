package queue

import (
	"context"
	"encoding/json"

	"github.com/HemlockPham7/common-libs/pkg/array"
	"github.com/newrelic/go-agent/v3/newrelic"
)

const BatchSize = 20

func (s *service) SendImportBookmarkJob(ctx context.Context, uid string, bookmarkInputs []*ImportBookmarkInput) error {
	span := newrelic.FromContext(ctx).StartSegment("SendImportBookmarkJob_QueueService")
	defer span.End()

	// split array into batches
	batches := array.SplitIntoBatches(bookmarkInputs, BatchSize)
	for _, batch := range batches {
		err := s.sendJob(ctx, uid, batch)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) sendJob(ctx context.Context, uid string, bookmarkInputs []*ImportBookmarkInput) error {
	span := newrelic.FromContext(ctx).StartSegment("SendImportBookmarkJobSendJob_QueueService")
	defer span.End()

	// create ImportMessage struct
	message := ImportMessage{
		UID:       uid,
		Bookmarks: bookmarkInputs,
	}

	// marshal ImportMessage struct to json
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// push message to redis queue
	return s.messageQueue.PushMessage(ctx, messageBytes)
}
