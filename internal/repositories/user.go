package repositories

import (
	"context"
	"go-clean-arch/internal/usecases"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Name      string
	Age       int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type userRepository struct {
	db gorm.DB
}

func NewUserRepository(db gorm.DB) usecases.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (s *userRepository) Create(ctx context.Context) error {
	return nil
}
