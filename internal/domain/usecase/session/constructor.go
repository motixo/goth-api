package session

import (
	"github.com/motixo/goth-api/internal/domain/service"
)

type SessionUseCase struct {
	sessionRepo Repository
	logger      service.Logger
	ulidGen     *service.ULIDGenerator
}

func NewUsecase(
	r Repository,
	logger service.Logger,
	ulidGen *service.ULIDGenerator,
) UseCase {
	return &SessionUseCase{
		sessionRepo: r,
		ulidGen:     ulidGen,
		logger:      logger,
	}
}
