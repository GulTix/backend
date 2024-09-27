package answers

import (
	"backend/internal/entity"
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *serviceImpl) Create(ctx context.Context, body CreateBody) (*CreateResponse, error) {
	var (
		answer *entity.Answer
		user   *entity.User
		err    error
	)

	// Find User by Email
	user, err = s.userRepo.FindUnique(ctx, body.Email, "email")

	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = s.userRepo.Create(
			ctx,
			entity.User{
				Id:          uuid.NewString(),
				FirstName:   body.FirstName,
				LastName:    body.LastName,
				PhoneNumber: body.PhoneNumber,
				Email:       body.Email,
				Gender:      body.Gender,
				Deleted:     false,
			},
		)

		if err != nil {
			return nil, err
		}
	}

	// Find Existing Answer
	answer, err = s.repo.FindExisting(ctx, user.Id, body.EventId)

	if err != nil {
		return nil, err
	}

	if answer != nil {
		err = fmt.Errorf("duplicate data")
		return nil, err
	}

	answer, err = s.repo.Create(
		ctx,
		entity.Answer{
			Id:      uuid.NewString(),
			EventId: body.EventId,
			UserId:  user.Id,
			RawData: body.RawData,
			Deleted: false,
		},
	)

	if err != nil {
		return nil, err
	}

	return &CreateResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Answer berhasil dibuat",
		Data:       *answer,
	}, nil
}
