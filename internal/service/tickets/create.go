package tickets

import (
	"backend/internal/entity"
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (s *serviceImpl) CreateResponse(ctx context.Context, data CreateBody) (*CreateResponse, error) {
	ticketTypeRaw := entity.TicketType{
		Id:       uuid.NewString(),
		Name:     data.Name,
		EventId:  data.EventId,
		Price:    data.Price,
		Quota:    data.Quota,
		BevyLink: data.BevyLink,
		Deleted:  false,
	}

	ticketType, err := s.repo.Create(ctx, ticketTypeRaw)

	if err != nil {
		return nil, err
	}

	return &CreateResponse{
		Success:    true,
		Message:    "Ticket Type berhasil dibuat",
		Data:       *ticketType,
		StatusCode: http.StatusCreated,
	}, nil
}
