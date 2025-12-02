package user

import (
	"context"

	"github.com/motixo/goth-api/internal/domain/entity"
)

type UseCase interface {
	GetProfile(ctx context.Context, userID string) (*entity.User, error)
	GetUser(ctx context.Context, userID string) (*entity.User, error)
}

type Repository interface {
	Create(ctx context.Context, u *entity.User) error
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}
