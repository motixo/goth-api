package session

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	client *redis.Client
}

func NewRepository(client *redis.Client) *Repository {
	return &Repository{client: client}
}

func (r *Repository) key(sessionID string) string {
	return fmt.Sprintf("session:%s", sessionID)
}
