package session

import "context"

func (r *Repository) DeleteSession(ctx context.Context, sessionID string) error {
	key := r.key(sessionID)
	return r.client.Del(ctx, key).Err()
}
