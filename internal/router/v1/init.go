package v1

import (
	"backend/internal/handler"
	"net/http"
)

func InitV1Router(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("/api", h.GetHandler)

	return http.StripPrefix("/api", mux)
}
