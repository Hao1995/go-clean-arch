package repositories

import (
	"context"
	"go-clean-arch/internal/usecases"
)

type rankRepository struct {
}

func NewRankRepository() usecases.RankRepository {
	return &rankRepository{}
}

func (s *rankRepository) Init(ctx context.Context) error {
	return nil
}
