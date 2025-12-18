package entity

import (
	"time"

	"github.com/motixo/goat-api/internal/domain/valueobject"
)

type Permission struct {
	ID        string                 `json:"id" db:"id"`
	Role      valueobject.UserRole   `json:"role" db:"role"`
	Action    valueobject.Permission `json:"action" db:"action"`
	CreatedAt time.Time              `json:"created_at" db:"created_at"`
}
