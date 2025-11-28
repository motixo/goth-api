package session

import (
	"context"
	"fmt"
	"time"

	"github.com/mot0x0/gopi/internal/domain/entity"
	"github.com/redis/go-redis/v9"
)

var createSessionLua = redis.NewScript(`
    -- KEYS[1]: key
    -- ARGV: sequence of field/value pairs + ttl

    local key = KEYS[1]
    local ttl = ARGV[#ARGV]   -- last argument is TTL
    local fieldCount = #ARGV - 1

    -- Build HSET arguments
    local hsetArgs = {}
    for i = 1, fieldCount do
        hsetArgs[i] = ARGV[i]
    end

    redis.call("HSET", key, unpack(hsetArgs))
    redis.call("EXPIRE", key, ttl)

    return 1
`)

func (r *Repository) CreateSession(ctx context.Context, s *entity.Session) error {
	key := r.key(s.ID)

	ttl := time.Until(s.ExpiresAt)
	if ttl <= 0 {
		return fmt.Errorf("expires_at is in the past")
	}

	argv := []interface{}{
		"user_id", s.UserID,
		"device", s.Device,
		"ip", s.IP,
		"created_at", s.CreatedAt.Unix(),
		"expires_at", s.ExpiresAt.Unix(),
		"current_jti", s.CurrentJTI,
		ttl.Seconds(),
	}

	_, err := createSessionLua.Run(ctx, r.client, []string{key}, argv...).Result()
	return err
}
