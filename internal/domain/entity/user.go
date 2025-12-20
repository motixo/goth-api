package entity

import (
	"time"

	"github.com/motixo/goat-api/internal/domain/valueobject"
)

type User struct {
	ID        string                 `db:"id"`
	Email     string                 `db:"email"`
	Password  valueobject.Password   `db:"password"`
	Status    valueobject.UserStatus `db:"status"`
	Role      valueobject.UserRole   `db:"role"`
	CreatedAt time.Time              `db:"created_at"`
	UpdatedAt *time.Time             `db:"updated_at"`
}
