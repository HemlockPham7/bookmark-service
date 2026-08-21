package queue

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (r *redisQueue) PushMessage(ctx context.Context, message []byte) error {
	span := newrelic.FromContext(ctx).StartSegment("PushMessage_QueueRepository")
	defer span.End()

	return r.client.LPush(ctx, r.queueName, message).Err()
}
