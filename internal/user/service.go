package user

import (
	"context"
	"errors"

	"github.com/mhafidk/warga-sekolah/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type Service struct {
	repo      *Repository
	jwtSecret []byte
}

func NewService(repo *Repository, jwtSecret []byte) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *Service) Authenticate(ctx context.Context, email, password string) (*User, string, error) {
	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", errors.New("Email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("Authentication failed")
	}

	token, err := auth.GenerateToken(u.ID, s.jwtSecret)
	if err != nil {
		return nil, "", err
	}

	return u, token, nil
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
		ID:       u.ID,
		Email:    email,
		FullName: fullName,
	}

	return u, nil
}
