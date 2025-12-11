package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/motixo/goat-api/internal/domain/entity"
	domainErrors "github.com/motixo/goat-api/internal/domain/errors"
	"github.com/motixo/goat-api/internal/domain/repository"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository.UserRepository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, u *entity.User) error {
	query := `
        INSERT INTO users (id, email, password, status, role, created_at, updated_at)
        VALUES (:id, :email, :password, :status, :role, :created_at, :updated_at)
    `
	_, err := r.db.NamedExecContext(ctx, query, u)
	return err
}

func (r *Repository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	query := `
        SELECT id, email, password, status, role, created_at, updated_at
        FROM users
        WHERE id = $1
		LIMIT 1
    `
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	query := `
        SELECT id, email, password, status, role, created_at, updated_at
        FROM users
        WHERE email = $1
		LIMIT 1
    `
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}

func (r *Repository) Update(ctx context.Context, user *entity.User) error {
	query := `
        UPDATE users 
        SET email = $1, password = $2, status = $3, role = $4, updated_at = $5 
        WHERE id = $6`

	result, err := r.db.ExecContext(ctx, query,
		user.Email,
		user.Password,
		user.Status,
		user.Role,
		time.Now(),
		user.ID,
	)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return domainErrors.ErrUserNotFound
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, userID string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", userID)
	return err
}

func (r *Repository) List(ctx context.Context, offset, limit int) ([]*entity.User, int64, error) {
	var users []*entity.User

	var total int64

	if err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, email, role, status, created_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	if err := r.db.SelectContext(ctx, &users, query, limit, offset); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
