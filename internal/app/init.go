package app

import (
	"backend/internal/handler"
	authHandler "backend/internal/handler/auth"
	answerService "backend/internal/service/answers"
	authService "backend/internal/service/auth"
	eventService "backend/internal/service/events"

	answerHandler "backend/internal/handler/answers"
	eventHandler "backend/internal/handler/events"
	formHandler "backend/internal/handler/forms"

	answerRepository "backend/internal/repository/answers"
	eventRepository "backend/internal/repository/events"
	userRepository "backend/internal/repository/users"

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
	eventRepo := eventRepository.NewRepository(db)
	answerRepo := answerRepository.NewRepository(db)

	authService := authService.NewService()
	oauthService := oauth.NewService(oauthClient, userRepo, authService)
	eventService := eventService.NewService(eventRepo)
	answerService := answerService.NewService(answerRepo, userRepo)

	authHandler := authHandler.NewHandler(oauthService)
	formHandler := formHandler.NewHandler()
	eventHandler := eventHandler.NewHandler(eventService)
	answerHandler := answerHandler.NewHandler(answerService)

	serverHandlers := handler.NewHandler(
		authHandler,
		formHandler,
		eventHandler,
		answerHandler,
	)
	server := NewServer(serverHandlers)
	return server
}
