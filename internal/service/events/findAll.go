package events

import (
	"backend/internal/entity"
	"context"
	"net/http"
)

func (s *serviceImpl) FindAll(ctx context.Context) (*FindAllResponse, error) {
	var (
		events     []entity.Event
		statusCode int = http.StatusOK
	)

	events, err := s.repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	if len(events) == 0 {
		statusCode = http.StatusNotFound
	}

	return &FindAllResponse{
		StatusCode: statusCode,
		Success:    true,
		Message:    "Ini data data event nya ya",
		Data:       events,
	}, nil

}
