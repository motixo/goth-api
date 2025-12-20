package entity

import (
	"time"

	"github.com/motixo/goat-api/internal/domain/valueobject"
)

type Permission struct {
	ID        string                 `db:"id"`
	Role      valueobject.UserRole   `db:"role"`
	Action    valueobject.Permission `db:"action"`
	CreatedAt time.Time              `db:"created_at"`
}
