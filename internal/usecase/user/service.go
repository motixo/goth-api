package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mot0x0/gopi/internal/domain/entities"
	"github.com/mot0x0/gopi/internal/domain/usecases"
	"github.com/mot0x0/gopi/internal/domain/valueobjects"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(ctx context.Context, u *entities.User) error
	GetByID(ctx context.Context, id string) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}

type UserUsecase struct {
	userRepo  UserRepository
	jwtSecret string
}

func NewUserUsecase(r UserRepository, secret string) usecases.UserUseCase {
	return &UserUsecase{
		userRepo:  r,
		jwtSecret: secret,
	}
}

// HashFuncs
func (u *UserUsecase) HashPassword(password string) (string, error) {
	salted := password + u.jwtSecret
	bytes, err := bcrypt.GenerateFromPassword([]byte(salted), bcrypt.DefaultCost)
	return string(bytes), err
}

// func CheckPassword(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

func (u *UserUsecase) Register(ctx context.Context, email string, password string) (*entities.User, error) {
	hashedPassword, err := u.HashPassword(password)
	if err != nil {
		return nil, err
	}

	rq := &entities.User{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  hashedPassword,
		Status:    valueobjects.StatusInactive,
		CreatedAt: time.Now().UTC(),
	}

	err = u.userRepo.Create(ctx, rq)
	if err != nil {
		return nil, err
	}
	return rq, nil
}

func (u *UserUsecase) Login(ctx context.Context, email, password string) (*entities.User, string, string, error) {
	// TODO: implement
	return nil, "", "", nil
}

func (u *UserUsecase) GetProfile(ctx context.Context, userID string) (*entities.User, error) {
	// TODO: implement
	return nil, nil
}

func (u *UserUsecase) ValidateToken(ctx context.Context, token string) (string, error) {
	// TODO: implement
	return "", nil
}
