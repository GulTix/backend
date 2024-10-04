package payments

import (
	"backend/internal/entity"
	"backend/pkg/database"
	"context"
)

type (
	Repository interface {
		Create(ctx context.Context, data entity.Payment) (*entity.Payment, error)
		UpdateStatus(ctx context.Context, paymentId string, status string, acquirer string) error
	}
	RepositoryImpl struct {
		db database.DB
	}
)
