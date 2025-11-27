package session

import "context"

func (s *SessionUseCase) Delete(ctx context.Context, sessionID string) error {
	return s.sessionRepo.DeleteSession(ctx, sessionID)
}
