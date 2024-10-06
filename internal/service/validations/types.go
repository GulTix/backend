package validations

import (
	"backend/internal/entity"
	"backend/internal/repository/validations"
	"backend/internal/service"
	"context"
)

type (
	Service interface {
		CreateResponse(ctx context.Context, body CreateBody) (*CreateResponse, error)
	}

	serviceImpl struct {
		repo validations.Repository
	}

	CreateResponse = service.BaseResponse[entity.Validation]

	CreateBody struct {
		AnswerId           string `json:"answer_id"`
		UserId             string `json:"user_id"`
		VolunteerId        string `json:"volunteer_id"`
		Classification     string `json:"classification"`
		IsSentPaymentEmail bool   `json:"is_sent_payment_email"`
	}
)
