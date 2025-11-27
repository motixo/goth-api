package session

type SessionUseCase struct {
	sessionRepo Repository
}

func NewSessionUsecase(r Repository) UseCase {
	return &SessionUseCase{
		sessionRepo: r,
	}
}
