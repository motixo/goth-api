package jti

type JTIUseCase struct {
	jtiRepo Repository
}

func NewJTIUsecase(r Repository) UseCase {
	return &JTIUseCase{
		jtiRepo: r,
	}
}
