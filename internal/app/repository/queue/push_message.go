package queue

import "context"

func (r *redisQueue) PushMessage(ctx context.Context, message []byte) error {
	return r.client.LPush(ctx, r.queueName, message).Err()
}
