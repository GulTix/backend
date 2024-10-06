package events

import (
	"backend/internal/entity"
	"backend/internal/repository/events"
	"backend/internal/service"
	"backend/pkg/oauth"
	"context"

	"golang.org/x/oauth2"
)

type (
	Service interface {
		Create(ctx context.Context, data CreateBody) (*CreateResponse, error)
		FindAll(ctx context.Context) (*FindAllResponse, error)
		SetGoogleTokenResponse(
			ctx context.Context,
			eventId string,
			code string,
		) (*SetGoogleTokenResponse, error)
		GetGoogleLoginURL() string
		GetBlasterToken(ctx context.Context, eventId string) (*oauth2.Token, error)
	}

	serviceImpl struct {
		repo  events.Repository
		oauth oauth.OAuth
	}

	CreateResponse         = service.BaseResponse[entity.Event]
	FindAllResponse        = service.BaseResponse[[]entity.Event]
	SetGoogleTokenResponse = service.BaseResponse[*entity.Event]

	CreateBody struct {
		Name           string `json:"name" db:"name"`
		GoogleFormLink string `json:"google_form_link" db:"google_form_link"`
	}
)
