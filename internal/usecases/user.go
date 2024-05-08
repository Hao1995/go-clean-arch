package usecases

import (
	"context"

	"github.com/pkg/errors"
)

type UserRepository interface {
	Create(ctx context.Context) error
}

type UserUseCase interface {
	Register(ctx context.Context) error
}

type userUseCase struct {
	unitOfWork     UnitOfWork
	userRepository UserRepository
	rankRepository RankRepository
}

func (s *userUseCase) Register(ctx context.Context) error {
	// Simulate operations across multiple repositories and services.
	return s.unitOfWork.BeginTx(ctx, func(tx UnitOfWork) error {
		if err := s.userRepository.Create(ctx); err != nil {
			return errors.Wrap(err, "failed to create user")
		}
		if err := s.rankRepository.Init(ctx); err != nil {
			return errors.Wrap(err, "failed to init rank")
		}
		return nil
	})
}
