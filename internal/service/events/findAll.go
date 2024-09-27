package events

import (
	"backend/internal/entity"
	"context"
)

func (s *serviceImpl) FindAll(ctx context.Context) (*FindAllResponse, error) {
	var (
		events []entity.Event
	)

	events, err := s.repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return &FindAllResponse{
		Success: true,
		Message: "Ini data data event nya ya",
		Data:    events,
	}, nil

}
