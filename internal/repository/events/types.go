package events

import (
	"backend/internal/entity"
	"backend/pkg/database"
	"context"

	"golang.org/x/oauth2"
)

type (
	Repository interface {
		Create(ctx context.Context, data entity.Event) (*entity.Event, error)
		FindAll(ctx context.Context) ([]entity.Event, error)
		UpdateToken(ctx context.Context, token []byte, eventId string) error
		GetBlasterToken(ctx context.Context, eventId string) (*oauth2.Token, error)
		Delete(ctx context.Context, eventId string) (*entity.Event, error)
		// FindAllByVolunteerId(ctx context.Context, volunteerId string) ([]entity.Event, error)
	}

	RepositoryImpl struct {
		db database.DB
	}
)
