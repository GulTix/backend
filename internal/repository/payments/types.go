package payments

import (
	"backend/internal/entity"
	"backend/pkg/database"
	"context"
)

type (
	Repository interface {
		Create(ctx context.Context, data entity.Payment) (*entity.Payment, error)
		UpdateStatus(ctx context.Context, paymentId string, status string, payload map[string]any) error
		UpdateQrisUrl(ctx context.Context, qrisUrl string, paymentId string) (*entity.Payment, error)
	}
	RepositoryImpl struct {
		db database.DB
	}
)
