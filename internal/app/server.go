package app

import (
	_ "backend/docs"
	"fmt"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"

	"backend/internal/handler"
	"backend/internal/router"
	"log"
	"net/http"
)

func NewServer(handlers *handler.Handlers) *Server {
	return &Server{
		handler: handlers,
	}
}

func (s *Server) InitRouteAndServe() {
	// config := config.InitConfig()
	mux := http.NewServeMux()

	// mux.HandleFunc("/", s.handler.Auth.ReturnHelloWorld)

	mux.HandleFunc("/test", s.handler.Auth.ReturnHelloWorld)

	mux.Handle("/", router.InitRouter(s.handler))

	mux.Handle("/docs/*", httpSwagger.Handler())

	port := os.Getenv("PORT")

	log.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
