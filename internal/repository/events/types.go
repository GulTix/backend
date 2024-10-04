package events

import (
	"backend/internal/entity"
	"backend/pkg/database"
	"context"
)

type (
	Repository interface {
		Create(ctx context.Context, data entity.Event) (*entity.Event, error)
		FindAll(ctx context.Context) ([]entity.Event, error)
		UpdateToken(ctx context.Context, token []byte, eventId string) error
		// FindAllByVolunteerId(ctx context.Context, volunteerId string) ([]entity.Event, error)
	}

	RepositoryImpl struct {
		db database.DB
	}
)
