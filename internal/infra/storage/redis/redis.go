package redis

import (
	"context"
	"fmt"

	"github.com/motixo/goat-api/internal/config"
	"github.com/motixo/goat-api/internal/pkg"
	"github.com/redis/go-redis/v9"
)

func NewClient(cfg *config.Config, logger pkg.Logger) (*redis.Client, error) {
	rdb := redis.NewClient(cfg.RedisOptions())

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		logger.Error("failed to connect to Redis", "error", err)
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	logger.Info("Redis connected successfully", "addr", rdb.Options().Addr)
	return rdb, nil
}
