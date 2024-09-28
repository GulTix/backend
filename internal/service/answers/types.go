package answers

import (
	"backend/internal/entity"
	"backend/internal/repository/answers"
	"backend/internal/repository/users"
	"backend/internal/service"
	"context"
)

type (
	Service interface {
		Create(ctx context.Context, body CreateBody) (*CreateResponse, error)
	}

	serviceImpl struct {
		repo     answers.Repository
		userRepo users.Repository
	}

	CreateResponse  = service.BaseResponse[entity.Answer]
	FindAllResponse = service.BaseResponse[[]entity.Answer]

	CreateBody struct {
		EventId     string `json:"eventId"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		PhoneNumber string `json:"phoneNumber"`
		Gender      string `json:"gender"`
		Email       string `json:"email"`
		RawData     map[string]any
	}
)
