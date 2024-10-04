package tickets

import (
	"backend/internal/entity"
	tickettype "backend/internal/repository/tickets"
	"backend/internal/service"
	"context"
)

type (
	Service interface {
		CreateResponse(ctx context.Context, body CreateBody) (*CreateResponse, error)
		FindByEventIdResponse(ctx context.Context, eventId string) (*FindAllResponse, error)
	}

	serviceImpl struct {
		repo tickettype.Repository
	}

	CreateResponse  = service.BaseResponse[entity.TicketType]
	FindAllResponse = service.BaseResponse[[]entity.TicketType]

	CreateBody struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quota    int     `json:"quota"`
		EventId  string  `json:"eventId"`
		BevyLink string  `json:"bevyLink"`
	}
)
