package classifications

import (
	"backend/internal/entity"
	"backend/internal/repository/allowedclassifications"
	"backend/internal/repository/classifications"
	"backend/internal/service"
	"context"
)

type (
	Service interface {
		Create(ctx context.Context, body CreateBody) (*CreateResponse, error)
		FindAll(ctx context.Context, eventId string) (*FindAllResponse, error)
		AppendAllowed(ctx context.Context, body AppendBody) (*AppendResponse, error)
	}
	ServiceImpl struct {
		classRepo      classifications.Repository
		allowClassRepo allowedclassifications.Repository
	}

	CreateResponse  = service.BaseResponse[entity.Classification]
	FindAllResponse = service.BaseResponse[[]entity.Classification]
	AppendResponse  = service.BaseResponse[any]

	CreateBody struct {
		Name     string `json:"name"`
		EventId  string `json:"event_id"`
		TicketId string `json:"ticket_id"`
	}

	AppendBody struct {
		ClassificationId string `json:"classification_id"`
		TicketId         string `json:"ticket_id"`
	}
)
