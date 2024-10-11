package forms

import (
	"backend/internal/handler"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /answer", h.Answer.Create)
	// mux.HandleFunc("GET /events/{id}", h.Event.)
	return http.StripPrefix("/forms", mux)
}
