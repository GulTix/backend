package classifications

import (
	"backend/internal/entity"
	"backend/pkg/database"
	"context"
)

type (
	Repository interface {
		Create(ctx context.Context, data entity.Classification) (*entity.Classification, error)
		FindAll(ctx context.Context) ([]entity.Classification, error)
		FindAllByEventId(ctx context.Context, eventId string) ([]entity.Classification, error)
	}
	RepositoryImpl struct {
		db database.DB
	}
)
