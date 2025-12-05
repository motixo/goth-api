package session

import (
	"github.com/motixo/goth-api/internal/domain/repository"
	"github.com/motixo/goth-api/internal/domain/service"
)

type SessionUseCase struct {
	sessionRepo repository.SessionRepository
	logger      service.Logger
	ulidGen     *service.ULIDGenerator
}

func NewUsecase(
	r repository.SessionRepository,
	logger service.Logger,
	ulidGen *service.ULIDGenerator,
) UseCase {
	return &SessionUseCase{
		sessionRepo: r,
		ulidGen:     ulidGen,
		logger:      logger,
	}
}
