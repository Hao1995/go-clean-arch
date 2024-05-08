package usecases

import "context"

type EmailService interface {
	Send(ctx context.Context) error
}
