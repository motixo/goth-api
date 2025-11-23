package repositories

import "github.com/mot0x0/gopi/internal/domain/entities"

type SessionRepository interface {
	Create(session *entities.Session) error
	FindByToken(token string) (*entities.Session, error)
	FindByUserID(userID string) (*entities.Session, error)
	Delete(token string) error
	DeleteByUserID(userID string) error
}
