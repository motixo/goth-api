package jti

import (
	"context"
)

type UseCase interface {
	StoreJTI(ctx context.Context, input StoreInput) error
	RevokeJTI(ctx context.Context, jti string) error
	IsJTIValid(ctx context.Context, jti string) (bool, error)
}

type Repository interface {
	SaveJTI(ctx context.Context, userID string, tokenID string, ttlSeconds int) error
	Exists(ctx context.Context, tokenID string) (bool, error)
	DeleteJTI(ctx context.Context, tokenID string) error
}
