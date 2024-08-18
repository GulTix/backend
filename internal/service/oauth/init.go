package oauth

import (
	"backend/internal/repository/user"
	"backend/internal/service/auth"
	"backend/pkg/oauth"
)

func NewService(
	oauth oauth.OAuth,
	userRepo user.Repository,
	authService auth.Service,
) Service {
	return &ServiceImpl{
		oauth:    oauth,
		userRepo: userRepo,
		authService: authService,
	}
}
