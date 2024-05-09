package repositories

import (
	"context"
	"go-clean-arch/internal/usecases"

	"gorm.io/gorm"
)

type rankRepository struct {
	db *gorm.DB
}

func NewRankRepository(db *gorm.DB) usecases.RankRepository {
	return &rankRepository{
		db: db,
	}
}

func (s *rankRepository) Init(ctx context.Context) error {
	return nil
}
