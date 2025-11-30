package session

import "context"

func (s *SessionUseCase) IsJTIValid(ctx context.Context, jti string) (bool, error) {
	valid, err := s.sessionRepo.ExistsJTI(ctx, jti)
	if err != nil {
		s.logger.Error("failed to check JTI validity", "jti", jti, "error", err)
		return false, err
	}
	s.logger.Debug("JTI validation result", "jti", jti, "valid", valid)
	return valid, nil
}
