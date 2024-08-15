package app

import (
	"backend/internal/handler"
	"backend/internal/handler/auth"
	oauth "backend/internal/service/oauth"
	oauth_pkg "backend/pkg/oauth"

	"github.com/joho/godotenv"
)

func InitHttp() *Server {
	_ = godotenv.Load(".env")

	oauthClient := oauth_pkg.NewClient()

	oauthService := oauth.NewService(oauthClient)

	authHandler := auth.NewHandler(oauthService)

	serverHandlers := handler.NewHandler(authHandler)
	server := NewServer(serverHandlers)
	return server
}
