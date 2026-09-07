package queue

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// PushMessage pushes a message to the Redis queue.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - message: the message payload to push to the queue.
//
// Returns:
//   - An error if the message cannot be pushed to the Redis queue.
func (r *redisQueue) PushMessage(ctx context.Context, message []byte) error {
	span := newrelic.FromContext(ctx).StartSegment("PushMessage_QueueRepository")
	defer span.End()

	return r.client.LPush(ctx, r.queueName, message).Err()
}
