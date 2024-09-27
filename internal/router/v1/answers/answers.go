package answers

import (
	"backend/internal/handler"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", h.Answer.Create)
	return http.StripPrefix("/answers", mux)
}
