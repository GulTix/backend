package forms

import (
	"backend/internal/handler"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Form.GetForms)
	return http.StripPrefix("/forms", mux)
}
