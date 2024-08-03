package app

import (
	"backend/internal/handler"

	"github.com/joho/godotenv"
)

func InitHttp() *Server {
	_ = godotenv.Load(".env")
	serverHandlers := handler.NewHandler()
	server := NewServer(serverHandlers)
	return server
}
