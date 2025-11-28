package session

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var updateSessionJTILua = redis.NewScript(`
    -- KEYS[1]: key
    -- ARGV[1]: jti
    -- ARGV[2]: ttl (seconds)

    local key = KEYS[1]
    local jti = ARGV[1]
    local ttl = ARGV[2]

    redis.call("HSET", key, "current_jti", jti)
    redis.call("EXPIRE", key, ttl)

    return 1
`)

func (r *Repository) UpdateSessionJTI(ctx context.Context, sessionID, jti string, ttlSeconds int) error {
	key := r.key(sessionID)

	_, err := updateSessionJTILua.Run(
		ctx,
		r.client,
		[]string{key},
		jti,
		ttlSeconds,
	).Result()

	return err
}
