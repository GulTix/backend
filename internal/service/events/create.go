package events

import (
	"backend/internal/entity"
	"context"
	"log"
	"net/http"

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

	log.Print("Masuk")

	return &CreateResponse{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "Event berhasil dibuat",
		Data:       *event,
	}, nil
}
