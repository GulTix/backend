package auth

import (
	"backend/internal/handler"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/google/login", h.Auth.GoogleLogin)
	return http.StripPrefix("/auth", mux)
}
