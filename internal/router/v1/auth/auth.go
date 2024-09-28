package auth

import (
	"backend/internal/handler"
	"net/http"
	"os"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/google/login", h.Auth.GoogleLogin)
	mux.HandleFunc("/google/callback", h.Auth.GoogleCallback)
	if os.Getenv("ENV") == "dev" {
		mux.HandleFunc("POST /debug/login", h.Auth.DevLogin)
	}
	return http.StripPrefix("/auth", mux)
}
