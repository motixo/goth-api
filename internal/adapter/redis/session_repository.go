package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type SessionRepo struct {
	client *redis.Client
}

func NewSessionRepository(client *redis.Client) *SessionRepo {
	return &SessionRepo{client: client}
}

func (r *SessionRepo) key(sessionID string) string {
	return fmt.Sprintf("session:%s", sessionID)
}
