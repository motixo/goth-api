package session

type SessionUseCase struct {
	sessionRepo Repository
}

func NewUsecase(r Repository) UseCase {
	return &SessionUseCase{
		sessionRepo: r,
	}
}
