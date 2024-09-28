package auth

import (
	"backend/internal/entity"
	"backend/internal/repository/volunteers"
	"backend/internal/service"
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Service interface {
		GenerateToken(volunteer entity.Volunteer) (*string, error)
		ReturnDevToken(ctx context.Context, email string) (*DebugTokenResponse, error)
	}

	serviceImpl struct {
		volunteerRepo volunteers.Repository
	}

	DebugTokenResponse = service.BaseResponse[LoginData]

	LoginData struct {
		Token string           `json:"token"`
		User  entity.Volunteer `json:"user"`
	}

	TokenClaims struct {
		jwt.Claims
		User entity.Volunteer
	}
)
