package service

import (
	"context"

	"github.com/motixo/goat-api/internal/domain/entity"
	"github.com/motixo/goat-api/internal/domain/valueobject"
)

type PermCacheService interface {
	GetRolePermissions(ctx context.Context, roleID valueobject.UserRole) ([]*entity.Permission, error)
	ClearCache(ctx context.Context, roleID valueobject.UserRole) error
}
