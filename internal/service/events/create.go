package events

import (
	"backend/internal/entity"
	"context"

	"github.com/google/uuid"
)

func (s *serviceImpl) Create(ctx context.Context, body CreateBody) (*CreateResponse, error) {
	var (
		event *entity.Event
	)

	eventRaw := entity.Event{
		Id:             uuid.NewString(),
		Name:           body.Name,
		GoogleFormLink: body.GoogleFormLink,
		Deleted:        false,
	}

	event, err := s.repo.Create(ctx, eventRaw)

	if err != nil {
		return nil, err
	}

	return &CreateResponse{
		Success: true,
		Message: "Event berhasil dibuat",
		Data:    *event,
	}, nil
}
