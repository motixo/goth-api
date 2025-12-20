package service

import (
	"context"
	"time"
)

type RateLimiter interface {
	Allow(
		ctx context.Context,
		actorType string,
		actorID string,
		resource string,
		limit int,
		windowDuration time.Duration,
	) (bool, time.Duration, int64, error)
}
