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

	message := "Classification created successfully"

	if data.TicketId != "" {
		err = s.allowClassRepo.Create(ctx, class.Id, data.TicketId)
		if err != nil {
			return nil, err
		}

		message += " and added to allowed ticket type"
	}

	return &CreateResponse{
		StatusCode: http.StatusCreated,
		Message:    message,
		Data:       *class,
	}, nil
}
