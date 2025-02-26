package router

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	v1 "backend/internal/router/v1"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	// v1Router := v1.InitV1Router(h)

	mux.HandleFunc("/", h.Auth.ReturnHelloWorld)
	mux.Handle("/v1/", v1.InitV1Router(h))

	muxJwt := middleware.JwtValidator(mux)
	muxCors := middleware.CORS(muxJwt)

	return middleware.Logger(muxCors)
}
