package answers

import (
	"backend/internal/handler"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", h.Answer.GetAll)
	return http.StripPrefix("/answers", mux)
}
