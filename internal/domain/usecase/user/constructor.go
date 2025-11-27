package user

type UserUseCase struct {
	userRepo Repository
}

func NewUserUsecase(r Repository) UseCase {
	return &UserUseCase{
		userRepo: r,
	}
}
