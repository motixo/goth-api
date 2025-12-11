package user

import (
	"context"

	"github.com/motixo/goat-api/internal/domain/entity"
	"github.com/motixo/goat-api/internal/domain/errors"
	"github.com/motixo/goat-api/internal/domain/repository"
	"github.com/motixo/goat-api/internal/domain/service"
)

type UserUseCase struct {
	userRepo       repository.UserRepository
	passwordHasher service.PasswordHasher
	userCache      service.UserCacheService
	sessionRepo    repository.SessionRepository
	logger         service.Logger
}

func NewUsecase(
	r repository.UserRepository,
	passwordHasher service.PasswordHasher,
	logger service.Logger,
	sessionRepo repository.SessionRepository,
	userCache service.UserCacheService,
) UseCase {
	return &UserUseCase{
		userRepo:       r,
		passwordHasher: passwordHasher,
		sessionRepo:    sessionRepo,
		userCache:      userCache,
		logger:         logger,
	}
}

func (us *UserUseCase) GetUser(ctx context.Context, userID string) (*UserResponse, error) {
	us.logger.Info("Fetching user by ID", "userID:", userID)
	user, err := us.userRepo.FindByID(ctx, userID)
	if err != nil {
		us.logger.Error("Failed to fetch user", "userID", userID, "error", err)
		return nil, err
	}
	response := &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role.String(),
		Status:    user.Status.String(),
		CreatedAt: user.CreatedAt,
	}
	us.logger.Info("User fetched successfully", "userID:", userID)
	return response, nil
}

func (us *UserUseCase) GetUserslist(ctx context.Context, offset, limit int) ([]*UserResponse, int64, error) {
	us.logger.Info("Fetching users List")
	users, total, err := us.userRepo.List(ctx, offset, limit)
	if err != nil {
		us.logger.Error("Failed to fetch users List", "error", err)
		return nil, 0, err
	}

	response := make([]*UserResponse, 0, len(users))
	for _, usr := range users {
		r := &UserResponse{
			ID:        usr.ID,
			Email:     usr.Email,
			Role:      usr.Role.String(),
			Status:    usr.Status.String(),
			CreatedAt: usr.CreatedAt,
		}
		response = append(response, r)
	}
	us.logger.Info("Users list fetched successfully")
	return response, total, nil
}

func (us *UserUseCase) DeleteUser(ctx context.Context, userID string) error {
	us.logger.Info("Attempting to delete user", "TargetUserID:", userID)
	if err := us.userRepo.Delete(ctx, userID); err != nil {
		us.logger.Error("Failed to delete user", "Error:", err)
		return err
	}

	sessions, _, err := us.sessionRepo.ListByUser(ctx, userID, 0, 0)
	if err != nil {
		us.logger.Error("field to fetch user sessions", "UserID:", userID)
		return nil
	}

	if len(sessions) == 0 {
		return nil
	}

	targets := make([]string, 0, len(sessions))
	for i := range sessions {
		targets[i] = sessions[i].ID
	}
	if err := us.sessionRepo.Delete(ctx, targets); err != nil {
		us.logger.Error("filed to delete user sessions", "userID:", userID)
		return nil
	}

	us.logger.Info("User deleted successfully", "TargetUserID:", userID)
	return nil
}

func (us *UserUseCase) ChangeEmail(ctx context.Context, input UpdateEmailInput) error {
	us.logger.Info("update user attempt", "UserID:", input.UserID)

	usr := &entity.User{
		ID:    input.UserID,
		Email: input.Email,
	}

	if err := us.userRepo.Update(ctx, usr); err != nil {
		us.logger.Error("user update failed", "UserID:", input.UserID)
		return err
	}

	us.logger.Info("user successfully updated", "UserID:", input.UserID)
	return nil
}

func (us *UserUseCase) ChangePassword(ctx context.Context, input UpdatePassInput) error {
	us.logger.Info("change password attempt", "UserID:", input.UserID)
	if input.OldPassword == input.NewPassword {
		us.logger.Error("passwords are same", "UserID:", input.UserID)
		return errors.ErrPasswordSameAsCurrent
	}

	user, err := us.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		us.logger.Error("user lookup failed", "UserID:", input.UserID, "Error:", err)
		return errors.ErrUserNotFound
	}

	if !us.passwordHasher.Verify(ctx, input.OldPassword, user.Password) {
		return errors.ErrInvalidPassword
	}

	hashedPassword, err := us.passwordHasher.Hash(ctx, input.NewPassword)
	if err != nil {
		us.logger.Error("password hashing failed", "UserID:", input.UserID, "Error:", err)
		return err
	}

	usr := &entity.User{
		ID:       user.ID,
		Password: hashedPassword,
	}
	if err := us.userRepo.Update(ctx, usr); err != nil {
		us.logger.Error("user update failed", "UserID:", input.UserID, "Error:", err)
		return err
	}

	sessions, _, err := us.sessionRepo.ListByUser(ctx, user.ID, 0, 0)
	if err != nil {
		us.logger.Error("field to fetch user sessions", "UserID:", input.UserID, "Error:", err)
		return nil
	}

	if len(sessions) == 0 {
		return nil
	}

	targets := make([]string, len(sessions))
	for i := range sessions {
		targets[i] = sessions[i].ID
	}
	if err := us.sessionRepo.Delete(ctx, targets); err != nil {
		us.logger.Error("filed to delete user sessions", "userID:", input.UserID, "Error:", err)
		return nil
	}

	us.logger.Info("password updated and sessions removed", "UserID:", input.UserID)
	return nil

}

func (us *UserUseCase) ChangeRole(ctx context.Context, input UpdateRoleInput) error {
	us.logger.Info("change role attempt", "UserID:", input.UserID)
	usr := &entity.User{
		ID:   input.UserID,
		Role: input.Role,
	}
	if err := us.userRepo.Update(ctx, usr); err != nil {
		us.logger.Error("change user role faild", "user_id", input.UserID)
		return err
	}
	if err := us.userCache.ClearCache(ctx, input.UserID); err != nil {
		us.logger.Error("clear user cache faild", "user_id", input.UserID)
	}
	us.logger.Info("user role changed successfully", "UserID:", input.UserID)
	return nil
}

func (us *UserUseCase) ChangeStatus(ctx context.Context, input UpdateStatusInput) error {
	us.logger.Info("change status attempt", "UserID:", input.UserID)
	usr := &entity.User{
		ID:     input.UserID,
		Status: input.Status,
	}
	if err := us.userRepo.Update(ctx, usr); err != nil {
		us.logger.Error("change user status faild", "user_id", input.UserID)
		return err
	}
	if err := us.userCache.ClearCache(ctx, input.UserID); err != nil {
		us.logger.Error("clear user cache faild", "user_id", input.UserID)
	}
	us.logger.Info("user status changed successfully", "UserID:", input.UserID)
	return nil
}
