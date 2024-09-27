package answers

import (
	"backend/internal/entity"
	"backend/pkg/database"
	"context"
)

type (
	Repository interface {
		Create(ctx context.Context, data entity.Answer) (*entity.Answer, error)
		FindAll(ctx context.Context) ([]entity.Answer, error)
		FindAllByEventId(ctx context.Context, eventId string) ([]entity.Answer, error)
	}

	RepositoryImpl struct {
		db database.DB
	}
)
