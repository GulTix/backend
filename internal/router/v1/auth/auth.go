package auth

import (
	"backend/internal/handler"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/google/login", h.Auth.GoogleLogin)
	mux.HandleFunc("/google/callback", h.Auth.GoogleCallback)
	return http.StripPrefix("/auth", mux)
}
