package events

import (
	"backend/internal/handler"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /token/login", h.Event.GoogleLogin)
	mux.HandleFunc("PUT /token/callback", h.Event.SetGoogleToken)
	mux.HandleFunc("GET /{id}/tickets/", h.Event.FindAllTicket)
	mux.HandleFunc("POST /{id}/tickets/", h.Event.CreateTicket)

	return http.StripPrefix("/events", mux)
}
