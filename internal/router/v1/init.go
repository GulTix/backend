package v1

import (
	"backend/internal/handler"
	"backend/internal/router/v1/auth"
	"net/http"
)

func InitV1Router(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", h.Auth.ReturnHelloWorld)
	mux.Handle("/auth/", auth.InitRouter(h))

	return http.StripPrefix("/v1", mux)
}
