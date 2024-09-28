package oauth

import (
	"backend/internal/repository/users"
	"backend/internal/repository/volunteers"
	"backend/internal/service/auth"
	"backend/pkg/oauth"
)

func NewService(
	oauth oauth.OAuth,
	userRepo users.Repository,
	volunteerRepo volunteers.Repository,
	authService auth.Service,
) Service {
	return &ServiceImpl{
		oauth:         oauth,
		userRepo:      userRepo,
		volunteerRepo: volunteerRepo,
		authService:   authService,
	}
}
