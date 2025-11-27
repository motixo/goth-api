package session

import (
	"context"

	"github.com/mot0x0/gopi/internal/domain/entity"
)

func (s *SessionUseCase) Create(ctx context.Context, session *entity.Session) error {
	return s.sessionRepo.CreateSession(ctx, session)
}
