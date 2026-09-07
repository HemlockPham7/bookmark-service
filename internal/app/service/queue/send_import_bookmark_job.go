package queue

import (
	"context"
	"encoding/json"

	"github.com/HemlockPham7/common-libs/pkg/array"
	"github.com/newrelic/go-agent/v3/newrelic"
)

const BatchSize = 20

// SendImportBookmarkJob sends bookmark import jobs to the message queue in batches.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - uid: the ID of the user who owns the bookmarks.
//   - bookmarkInputs: the bookmark data to be imported.
//
// Returns:
//   - An error if any batch cannot be sent to the message queue.
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

// sendJob creates and queues an import bookmark job for a batch of bookmarks.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - uid: the ID of the user who owns the bookmarks.
//   - bookmarkInputs: the bookmark data included in the job batch.
//
// Returns:
//   - An error if the job cannot be serialized or pushed to the message queue.
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
