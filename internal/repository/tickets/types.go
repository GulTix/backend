package tickets

import (
	"backend/internal/entity"
	"backend/pkg/database"
	"context"
)

type (
	Repository interface {
		Create(ctx context.Context, data entity.TicketType) (*entity.TicketType, error)
		FindAllByEventId(ctx context.Context, eventId string) ([]entity.TicketType, error)
	}

	RepositoryImpl struct {
		db database.DB
	}
)
