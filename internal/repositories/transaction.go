package repositories

import (
	"context"
	"fmt"
	"go-clean-arch/internal/usecases"

	"gorm.io/gorm"
)

type unitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) usecases.UnitOfWork {
	return &unitOfWork{
		db: db,
	}
}

func (s *unitOfWork) BeginTx(ctx context.Context, fn func(tx usecases.UnitOfWork) error) error {
	var err error
	tx := s.db.Begin()
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback().Error; rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
				fmt.Printf("%v\n", err)
			}
		} else {
			err = tx.Commit().Error
			fmt.Printf("%v\n", err)
		}
	}()

	newUnitOfWork := &unitOfWork{
		db: tx,
	}
	err = fn(newUnitOfWork)
	return nil
}

func (s *unitOfWork) NewUserRepository() usecases.UserRepository {
	return &userRepository{
		db: s.db,
	}
}

func (s *unitOfWork) NewRankRepository() usecases.RankRepository {
	return &rankRepository{
		db: s.db,
	}
}
