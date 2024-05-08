package repositories

import (
	"context"
	"go-clean-arch/internal/usecases"
)

type email struct {
}

func NewEmailService() usecases.EmailService {
	return &email{}
}

func (s *email) Send(ctx context.Context) error {
	return nil
}
