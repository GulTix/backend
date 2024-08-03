package router

import (
	"backend/internal/handler"
	v1 "backend/internal/router/v1"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	v1Router := v1.InitV1Router(h)

	mux.Handle("/v1", v1Router)

	return http.StripPrefix("/api", mux)
}
