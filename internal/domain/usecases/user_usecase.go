package usecases

import (
	"context"

	"github.com/mot0x0/gopi/internal/domain/entities"
)

type UserUseCase interface {
	Register(ctx context.Context, email, password string) (*entities.User, error)
	Login(ctx context.Context, email, password string) (*entities.User, string, string, error)
	GetProfile(ctx context.Context, userID string) (*entities.User, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}
