package events

import (
	"backend/internal/entity"
	"backend/internal/repository/events"
	"backend/internal/service"
	"context"
)

type (
	Service interface {
		Create(ctx context.Context, data CreateBody) (*CreateResponse, error)
		FindAll(ctx context.Context) (*FindAllResponse, error)
	}

	serviceImpl struct {
		repo events.Repository
	}

	CreateResponse  = service.BaseResponse[entity.Event]
	FindAllResponse = service.BaseResponse[[]entity.Event]

	CreateBody struct {
		Name           string `json:"name" db:"name"`
		BevyLink       string `json:"bevy_link" db:"bevy_link"`
		GoogleFormLink string `json:"google_form_link" db:"google_form_link"`
	}
)
