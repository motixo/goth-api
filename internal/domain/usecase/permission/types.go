package permission

import (
	"time"

	"github.com/motixo/goat-api/internal/domain/valueobject"
)

type CreateInput struct {
	Role   valueobject.UserRole   `json:"role"`
	Action valueobject.Permission `json:"action"`
}

type PermissionResponse struct {
	ID        string    `json:"id"`
	Role      string    `json:"role"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
}
