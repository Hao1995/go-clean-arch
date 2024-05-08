package usecases

import "context"

type RankRepository interface {
	Init(ctx context.Context) error
}
