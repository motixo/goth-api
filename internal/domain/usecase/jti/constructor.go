package jti

type JTIUseCase struct {
	jtiRepo Repository
}

func NewUsecase(r Repository) UseCase {
	return &JTIUseCase{
		jtiRepo: r,
	}
}
