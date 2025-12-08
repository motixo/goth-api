package session

import (
	"github.com/motixo/goat-api/internal/domain/repository"
	"github.com/motixo/goat-api/internal/domain/service"
	"github.com/motixo/goat-api/internal/infra/logger"
)

type SessionUseCase struct {
	sessionRepo repository.SessionRepository
	logger      logger.Logger
}

func NewUsecase(
	r repository.SessionRepository,
	logger logger.Logger,
	ulidGen service.IDGenerator,
) UseCase {
	return &SessionUseCase{
		sessionRepo: r,
		logger:      logger,
	}
}
