package session

import "context"

func (us *SessionUseCase) IsJTIValid(ctx context.Context, jti string) (bool, error) {
	valid, err := us.sessionRepo.ExistsJTI(ctx, jti)
	if err != nil {
		us.logger.Error("failed to check JTI validity", "jti", jti, "error", err)
		return false, err
	}
	us.logger.Debug("JTI validation result", "jti", jti, "valid", valid)
	return valid, nil
}
