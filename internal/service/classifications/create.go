package classifications

import (
	"backend/internal/entity"
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (s *ServiceImpl) Create(ctx context.Context, data CreateBody) (*CreateResponse, error) {
	class, err := s.classRepo.Create(ctx, entity.Classification{
		Id:      uuid.NewString(),
		EventId: data.EventId,
		Name:    data.Name,
	})

	if err != nil {
		return nil, err
	}

	return &CreateResponse{
		StatusCode: http.StatusCreated,
		Message:    "Classification created successfully",
		Data:       *class,
	}, nil
}
