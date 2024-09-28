package auth

import (
	"context"
	"errors"
	"net/http"
)

func (s *serviceImpl) ReturnDevToken(
	ctx context.Context,
	email string,
) (
	*DebugTokenResponse,
	error,
) {
	volunteer, err := s.volunteerRepo.FindUnique(ctx, email, "email")

	if err != nil {
		return nil, err
	}

	if volunteer == nil {
		return nil, errors.New("volunteer not found")
	}

	token, err := s.GenerateToken(*volunteer)

	if err != nil {
		return nil, err
	}

	return &DebugTokenResponse{
		Data: LoginData{
			Token: *token,
			User:  *volunteer,
		},
		Success:    true,
		StatusCode: http.StatusOK,
	}, nil
}
