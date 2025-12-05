package permission

import (
	"github.com/motixo/goth-api/internal/domain/repository"
	"github.com/motixo/goth-api/internal/domain/service"
)

type PermissionUseCase struct {
	permissionRepo repository.PermissionRepository
	logger         service.Logger
}

func NewUsecase(
	p repository.PermissionRepository,
	logger service.Logger,
) UseCase {
	return &PermissionUseCase{
		permissionRepo: p,
		logger:         logger,
	}
}
