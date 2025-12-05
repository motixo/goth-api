package session

import (
	"context"
)

func (us *SessionUseCase) DeleteSessions(ctx context.Context, input DeleteSessionsInput) error {
	us.logger.Info("delete sessions requested", "userID", input.UserID, "removeOthers", input.RemoveOthers, "targetCount", len(input.TargetSessions))
	var target []string
	if input.RemoveOthers {
		sessions, err := us.GetSessionsByUser(ctx, input.UserID, input.CurrentSession)
		if err != nil {
			us.logger.Error("failed to get user sessions for deletion", "userID", input.UserID, "error", err)
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
		us.logger.Debug("no sessions to delete", "userID", input.UserID)
		return nil
	}

	err := us.sessionRepo.Delete(ctx, target)
	if err != nil {
		us.logger.Error("failed to delete sessions", "userID", input.UserID, "targetCount", len(target), "error", err)
		return err
	}
	us.logger.Info("sessions deleted successfully", "userID", input.UserID, "removeOthers", input.RemoveOthers, "targetCount", len(input.TargetSessions))
	return nil
}
