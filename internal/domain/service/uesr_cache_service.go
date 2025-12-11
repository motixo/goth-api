package service

import (
	"context"

	"github.com/motixo/goat-api/internal/domain/valueobject"
)

type UserCacheService interface {
	GetUserStatus(ctx context.Context, userID string) (valueobject.UserStatus, error)
	GetUserRole(ctx context.Context, userID string) (valueobject.UserRole, error)
	ClearCache(ctx context.Context, userID string) error
}
