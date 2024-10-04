package validations

import (
	"backend/internal/entity"
	"backend/pkg/database"
	"context"
)

type (
	Repository interface {
		Create(ctx context.Context, data entity.Validation) (*entity.Validation, error)
	}

	RepositoryImpl struct {
		db database.DB
	}
)
