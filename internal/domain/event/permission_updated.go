package event

import "github.com/motixo/goat-api/internal/domain/valueobject"

type PermissionUpdatedEvent struct {
	Role      valueobject.UserRole
	UpdatedBy string
}
