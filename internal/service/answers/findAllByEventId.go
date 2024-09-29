package answers

import (
	"backend/internal/entity"
	"context"
	"net/http"
)

func (s *serviceImpl) FindAllByEventId(ctx context.Context, eventId string) (*FindAllResponse, error) {
	var (
		answers    []entity.Answer
		statusCode int = http.StatusOK
	)

	answers, err := s.repo.FindAllByEventId(ctx, eventId)

	if err != nil {
		return nil, err
	}

	if len(answers) == 0 {
		statusCode = http.StatusNotFound
	}

	return &FindAllResponse{
		StatusCode: statusCode,
		Success:    true,
		Message:    "Ini data data answer nya ya",
		Data:       answers,
	}, nil
}
