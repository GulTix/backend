package validations

import (
	"backend/internal/entity"
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (s *serviceImpl) CreateResponse(ctx context.Context, body CreateBody) (*CreateResponse, error) {
	validationRaw := entity.Validation{
		Id:               uuid.NewString(),
		AnswerId:         body.AnswerId,
		UserId:           body.UserId,
		VolunteerId:      body.VolunteerId,
		Classification:   body.Classification,
		PaymentEmailSent: nil,
	}

	validation, err := s.repo.Create(ctx, validationRaw)

	if err != nil {
		return nil, err
	}

	return &CreateResponse{
		Success:    true,
		Message:    "Validation berhasil dibuat",
		Data:       *validation,
		StatusCode: http.StatusCreated,
	}, nil

}
