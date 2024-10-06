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
		FindAll(ctx context.Context) (*FindAllResponse, error)
	}
	ServiceImpl struct {
		classRepo      classifications.Repository
		allowClassRepo allowedclassifications.Repository
	}

	CreateResponse  = service.BaseResponse[entity.Classification]
	FindAllResponse = service.BaseResponse[[]entity.Classification]

	CreateBody struct {
		Name     string `json:"name"`
		EventId  string `json:"event_id"`
		TicketId string `json:"ticket_id"`
	}
)
