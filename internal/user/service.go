package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, email, fullName, password string) (*User, error) {
	exists, err := s.repo.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("Email is already registered.")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		Email:    email,
		FullName: fullName,
		Password: string(hashedPassword),
	}

	if err := s.repo.Save(ctx, u); err != nil {
		return nil, err
	}

	u = &User{
		Email:    email,
		FullName: fullName,
	}

	return u, nil
}
