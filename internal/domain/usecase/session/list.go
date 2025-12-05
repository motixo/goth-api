package session

import (
	"context"
	"sort"
)

func (us *SessionUseCase) GetSessionsByUser(ctx context.Context, userID, sessionID string) ([]*SessionResponse, error) {
	us.logger.Debug("retrieving user sessions", "userID", userID, "currentSessionID", sessionID)
	sessions, err := us.sessionRepo.ListByUser(ctx, userID)
	if err != nil {
		us.logger.Error("failed to list sessions by user", "userID", userID, "error", err)
		return []*SessionResponse{}, err
	}

	response := make([]*SessionResponse, 0, len(sessions))
	for _, se := range sessions {
		r := &SessionResponse{
			ID:        se.ID,
			Device:    se.Device,
			IP:        se.IP,
			Current:   se.ID == sessionID,
			CreatedAt: se.CreatedAt,
			UpdatedAt: se.UpdatedAt,
		}

		response = append(response, r)
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].UpdatedAt.After(response[j].UpdatedAt)
	})

	us.logger.Info("user sessions retrieved", "userID", userID, "sessionCount", len(sessions))
	return response, nil
}
