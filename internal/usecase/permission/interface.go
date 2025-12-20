package permission

import (
	"context"

	"github.com/motixo/goat-api/internal/domain/valueobject"
)

type UseCase interface {
	Create(ctx context.Context, input CreateInput) (PermissionResponse, error)
	GetPermissions(ctx context.Context, offset, limit int) ([]PermissionResponse, int64, error)
	GetPermissionsByRole(ctx context.Context, roleID valueobject.UserRole) ([]PermissionResponse, error)
	Delete(ctx context.Context, permissionID string) error
}
