package auth

import (
	"context"

	"github.com/motixo/goth-api/internal/domain/usecase/session"
)

func (a *AuthUseCase) Logout(ctx context.Context, sessionID, userID string) error {

	a.logger.Info("user logout requested", "userID", userID)

	input := session.DeleteSessionsInput{
		TargetSessions: []string{sessionID},
		UserID:         userID,
	}

	err := a.sessionUC.DeleteSessions(ctx, input)
	if err != nil {
		return err
	}
	a.logger.Info("user logged out", "userID", userID)
	return nil
}
