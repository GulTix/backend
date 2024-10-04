package tickets

import (
	"backend/internal/handler"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", h.Ticket.Create)

	return http.StripPrefix("/tickets", mux)
}
