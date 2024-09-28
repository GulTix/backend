package auth

import (
	"backend/internal/service/auth"
	"backend/internal/service/oauth"
)

func NewHandler(
	oauthService oauth.Service,
	authService auth.Service,
) Handler {
	return &HandlerImpl{
		authService:  authService,
		oauthService: oauthService,
	}
}
