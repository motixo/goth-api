package user

import (
	"context"

	"github.com/motixo/goth-api/internal/domain/entity"
)

func (u *UserUseCase) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	u.logger.Info("Fetching user by ID", "userID:", userID)
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		u.logger.Error("Failed to fetch user", "userID", userID, "error", err)
		return nil, err
	}
	u.logger.Info("User fetched successfully", "userID:", userID)
	return user, nil
}
