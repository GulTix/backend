package app

import (
	"backend/internal/handler"
	authHandler "backend/internal/handler/auth"
	answerService "backend/internal/service/answers"
	authService "backend/internal/service/auth"
	classificationsService "backend/internal/service/classifications"
	eventService "backend/internal/service/events"
	paymentsService "backend/internal/service/payments"
	ticketTypeService "backend/internal/service/tickets"
	validationService "backend/internal/service/validations"

	answerHandler "backend/internal/handler/answers"
	classHandler "backend/internal/handler/classifications"
	eventHandler "backend/internal/handler/events"
	formHandler "backend/internal/handler/forms"
	paymentHandler "backend/internal/handler/payments"
	ticketTypeHandler "backend/internal/handler/tickets"
	validationHandler "backend/internal/handler/validations"

	allowedClassificationRepository "backend/internal/repository/allowedclassifications"
	answerRepository "backend/internal/repository/answers"
	classificationRepository "backend/internal/repository/classifications"
	eventRepository "backend/internal/repository/events"
	paymentRepository "backend/internal/repository/payments"
	ticketTypeRepository "backend/internal/repository/tickets"
	userRepository "backend/internal/repository/users"
	validationRepository "backend/internal/repository/validations"
	volunteerRepository "backend/internal/repository/volunteers"

	oauth "backend/internal/service/oauth"
	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/midtrans"
	"backend/pkg/migrations"
	oauth_pkg "backend/pkg/oauth"
	"context"
	"os"

	"github.com/joho/godotenv"
)

func InitHttp() *Server {
	_ = godotenv.Load(".env")

	// Initialize Config
	config := config.InitConfig()

	// Running Migrations
	migrations.Migrate(config)

	ctx := context.Background()
	connetionString := config.DBConnectionString

	db := database.NewDB(ctx, connetionString)

	oauthLoginClient := oauth_pkg.NewClient([]string{
		"https://www.googleapis.com/auth/userinfo.email",
	})

	midtransClient := midtrans.NewMidtrans(config)

	oauthBlasterClient := oauth_pkg.NewClient([]string{
		"https://www.googleapis.com/auth/gmail.send",
	})

	userRepo := userRepository.NewRepository(db)
	eventRepo := eventRepository.NewRepository(db)
	answerRepo := answerRepository.NewRepository(db)
	volunteerRepo := volunteerRepository.NewRepository(db)
	validationRepo := validationRepository.NewRepository(db)
	ticketTypeRepository := ticketTypeRepository.NewRepository(db)
	paymentsRepo := paymentRepository.NewRepository(db)
	classificationRepo := classificationRepository.NewRepository(db)
	allowedClassRepo := allowedClassificationRepository.NewRepository(db)

	if os.Getenv("ENV") == "dev" {
		err := volunteerRepo.CreateDevVolunteer(ctx)
		if err != nil {
			panic(err)
		}
	}

	authService := authService.NewService(volunteerRepo)
	oauthService := oauth.NewService(
		oauthLoginClient,
		userRepo,
		volunteerRepo,
		authService,
	)
	eventService := eventService.NewService(
		eventRepo,
		oauthBlasterClient,
	)
	answerService := answerService.NewService(answerRepo, userRepo)
	validationService := validationService.NewService(validationRepo)
	ticketTypeService := ticketTypeService.NewService(ticketTypeRepository)
	paymentsService := paymentsService.NewService(midtransClient, paymentsRepo, eventService, oauthBlasterClient)
	classificationsService := classificationsService.NewService(classificationRepo, allowedClassRepo)

	authHandler := authHandler.NewHandler(
		oauthService,
		authService,
	)

	formHandler := formHandler.NewHandler()
	eventHandler := eventHandler.NewHandler(
		eventService,
		ticketTypeService,
		classificationsService,
	)

	answerHandler := answerHandler.NewHandler(answerService)
	validationHandler := validationHandler.NewHandler(validationService)
	ticketTypeHandler := ticketTypeHandler.NewHandler(ticketTypeService, classificationsService)
	paymentHandler := paymentHandler.NewHandler(paymentsService)
	classHandler := classHandler.NewHandler(classificationsService)

	serverHandlers := handler.NewHandler(
		authHandler,
		formHandler,
		eventHandler,
		answerHandler,
		ticketTypeHandler,
		validationHandler,
		paymentHandler,
		classHandler,
	)

	server := NewServer(serverHandlers)
	return server
}
