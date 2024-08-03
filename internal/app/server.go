package app

import (
	_ "backend/docs"
	"backend/pkg/migrations"

	httpSwagger "github.com/swaggo/http-swagger"

	"backend/internal/handler"
	"backend/internal/router"
	"log"
	"net/http"
	"os"
)

func NewServer(handlers *handler.Handlers) *Server {
	return &Server{
		handler: handlers,
	}
}

func (s *Server) InitRouteAndServe() {
	mux := http.NewServeMux()

	mux.Handle("/docs/*", httpSwagger.Handler())

	mux.Handle("/api", router.InitRouter(s.handler))

	migrations.Migrate()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(":"+port, mux)
}
