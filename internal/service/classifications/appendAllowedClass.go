package classifications

import (
	"context"
	"net/http"
)

func (s *ServiceImpl) AppendAllowed(ctx context.Context, data AppendBody) (*AppendResponse, error) {
	err := s.allowClassRepo.Create(ctx, data.ClassificationId, data.TicketId)

	if err != nil {
		return nil, err
	}

	return &AppendResponse{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "Berhasil di tambahkan",
		Data:       nil,
	}, nil
}
