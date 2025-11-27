package auth

import (
	"github.com/mot0x0/gopi/internal/domain/usecase/jti"
	"github.com/mot0x0/gopi/internal/domain/usecase/user"
)

type AuthUseCase struct {
	userRepo user.Repository
	jtiUC    jti.UseCase
}

func NewAuthUsecase(jtiUC jti.UseCase, userRepo user.Repository) UseCase {
	return &AuthUseCase{
		userRepo: userRepo,
		jtiUC:    jtiUC,
	}
}
