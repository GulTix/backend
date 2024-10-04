package tickets

import (
	"backend/internal/entity"
	"context"
	"net/http"
)

func (s *serviceImpl) FindByEventIdResponse(ctx context.Context, eventId string) (*FindAllResponse, error) {
	var tickets []entity.TicketType

	tickets, err := s.repo.FindAllByEventId(ctx, eventId)

	if err != nil {
		return nil, err
	}

	return &FindAllResponse{
		Success: true,
		Message: "Ini semua data tiketnya ya bwang",
		Data: tickets,
		StatusCode: http.StatusOK,
	}, nil
}
