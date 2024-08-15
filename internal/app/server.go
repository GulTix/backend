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

	migrations.Migrate()

	// mux.HandleFunc("/", s.handler.Auth.ReturnHelloWorld)

	mux.HandleFunc("/test", s.handler.Auth.ReturnHelloWorld)

	mux.Handle("/", router.InitRouter(s.handler))

	mux.Handle("/docs/*", httpSwagger.Handler())

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(":"+port, mux)
}
