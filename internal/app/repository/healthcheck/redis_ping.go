package healthcheck

import "context"

func (r *healthCheckRepository) RedisPing(ctx context.Context) error {
	return r.redisClient.Ping(ctx).Err()
}
