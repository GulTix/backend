package app

import (
	"backend/internal/handler"
	authHandler "backend/internal/handler/auth"
	authService "backend/internal/service/auth"

	formHandler "backend/internal/handler/forms"

	userRepository "backend/internal/repository/user"
	oauth "backend/internal/service/oauth"
	"backend/pkg/database"
	oauth_pkg "backend/pkg/oauth"
	"context"
	"os"

	"github.com/joho/godotenv"
)

func InitHttp() *Server {
	_ = godotenv.Load(".env")

	ctx := context.Background()
	connetionString := os.Getenv("DATABASE_URL")

	db := database.NewDB(ctx, connetionString)
	oauthClient := oauth_pkg.NewClient()

	userRepo := userRepository.NewRepository(db)

	authService := authService.NewService()
	oauthService := oauth.NewService(oauthClient, userRepo, authService)

	authHandler := authHandler.NewHandler(oauthService)
	formHandler := formHandler.NewHandler()

	serverHandlers := handler.NewHandler(authHandler, formHandler)
	server := NewServer(serverHandlers)
	return server
}
