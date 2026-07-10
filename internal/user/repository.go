package user

import (
	"context"
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByUserID(ctx context.Context, userID int64) (*User, error) {
	query := `SELECT id, full_name, email FROM users WHERE id = ? LIMIT 1`

	var u User
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&u.ID, &u.FullName, &u.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &u, nil
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `SELECT id, full_name, email, password FROM users WHERE email = ? LIMIT 1`

	var u User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.FullName, &u.Email, &u.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &u, nil
}

func (r *Repository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)`

	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
	return exists, err
}

func (r *Repository) Save(ctx context.Context, u *User) error {
	query := `INSERT INTO users (email, full_name, password) VALUES (?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, u.Email, u.FullName, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id
	return nil
}
