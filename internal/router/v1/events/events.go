package events

import (
	"backend/internal/handler"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", h.Event.Create)
	mux.HandleFunc("GET /", h.Event.FindAll)
	return http.StripPrefix("/events", mux)
}
