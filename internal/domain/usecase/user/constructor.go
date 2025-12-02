package user

import "github.com/motixo/goth-api/internal/domain/service"

type UserUseCase struct {
	userRepo Repository
	logger   service.Logger
}

func NewUsecase(r Repository, logger service.Logger) UseCase {
	return &UserUseCase{
		userRepo: r,
		logger:   logger,
	}
}
