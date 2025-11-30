package session

import (
	"context"
)

type DeleteSessionsInput struct {
	UserID         string
	CurrentSession string
	TargetSessions []string `json:"session_ids"`
	RemoveOthers   bool     `json:"others"`
}

func (u *SessionUseCase) DeleteSessions(ctx context.Context, input DeleteSessionsInput) error {
	u.logger.Info("delete sessions requested", "userID", input.UserID, "removeOthers", input.RemoveOthers, "targetCount", len(input.TargetSessions))
	var target []string
	if input.RemoveOthers {
		sessions, err := u.GetSessionsByUser(ctx, input.UserID, input.CurrentSession)
		if err != nil {
			u.logger.Error("failed to get user sessions for deletion", "userID", input.UserID, "error", err)
			return err
		}

		for _, s := range sessions {
			if !s.Current {
				target = append(target, s.ID)
			}
		}
	} else {
		target = input.TargetSessions
	}

	if len(target) == 0 {
		u.logger.Debug("no sessions to delete", "userID", input.UserID)
		return nil
	}

	err := u.sessionRepo.Delete(ctx, target)
	if err != nil {
		u.logger.Error("failed to delete sessions", "userID", input.UserID, "targetCount", len(target), "error", err)
		return err
	}
	u.logger.Info("sessions deleted successfully", "userID", input.UserID, "removeOthers", input.RemoveOthers, "targetCount", len(input.TargetSessions))
	return nil
}
