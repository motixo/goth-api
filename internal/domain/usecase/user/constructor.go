package user

type UserUseCase struct {
	userRepo Repository
}

func NewUsecase(r Repository) UseCase {
	return &UserUseCase{
		userRepo: r,
	}
}
