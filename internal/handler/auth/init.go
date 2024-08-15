package auth

import "backend/internal/service/oauth"

func NewHandler(oauthService oauth.Service) Handler {
	return &HandlerImpl{
		oauthService: oauthService,
	}
}
