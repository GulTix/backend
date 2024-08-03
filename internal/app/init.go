package app

import "backend/internal/handler"

func InitHttp() *Server {
	serverHandlers := handler.NewHandler()
	server := NewServer(serverHandlers)
	return server
}
