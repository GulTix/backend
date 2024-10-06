package classifications

import (
	"context"
	"net/http"
)

func (s *ServiceImpl) FindAll(ctx context.Context, eventId string) (*FindAllResponse, error) {
	classifications, err := s.classRepo.FindAllByEventId(ctx, eventId)

	if err != nil {
		return nil, err
	}

	return &FindAllResponse{
		StatusCode: http.StatusOK,
		Data:       classifications,
		Message:    "Success get all classifications",
		Success:    true,
	}, nil
}
