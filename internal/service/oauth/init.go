package oauth

import (
	"backend/internal/repository/users"
	"backend/internal/service/auth"
	"backend/pkg/oauth"
)

func NewService(
	oauth oauth.OAuth,
	userRepo users.Repository,
	authService auth.Service,
) Service {
	return &ServiceImpl{
		oauth:       oauth,
		userRepo:    userRepo,
		authService: authService,
	}
}
