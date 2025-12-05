package repository

import (
	"context"

	"github.com/motixo/goth-api/internal/domain/entity"
)

type PermissionRepository interface {
	Create(ctx context.Context, p *entity.Permission) error
	GetByRoleID(ctx context.Context, roleID int8) (*[]entity.Permission, error)
}
