package oauth

import (
	"backend/internal/entity"
	"backend/internal/repository/users"
	"backend/internal/repository/volunteers"
	"backend/internal/service"
	"backend/internal/service/auth"
	"backend/pkg/oauth"
	"context"
)

type (
	Service interface {
		GetGoogleLoginURL() string
		ReturnGoogleCallbackResponse(ctx context.Context, code string) (*GoogleCallbackResponse, error)
		// GoogleLogin(ctx context.Context) error
		// GoogleCallback(ctx context.Context) error
	}

	ServiceImpl struct {
		oauth         oauth.OAuth
		userRepo      users.Repository
		volunteerRepo volunteers.Repository
		authService   auth.Service
	}

	GoogleCallbackResponse = service.BaseResponse[LoginData]

	LoginData struct {
		Token string           `json:"token"`
		User  entity.Volunteer `json:"user"`
	}

	GoogleUserData struct {
		Sub           string `json:"sub"`
		Picture       string `json:"picture"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
	}
)
