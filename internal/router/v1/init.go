package v1

import (
	"backend/internal/handler"
	"backend/internal/router/v1/answers"
	"backend/internal/router/v1/auth"
	"backend/internal/router/v1/events"
	"backend/internal/router/v1/forms"
	"net/http"
)

func InitV1Router(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Auth.ReturnHelloWorld)
	mux.Handle("/auth/", auth.InitRouter(h))
	mux.Handle("/forms/", forms.InitRouter(h))
	mux.Handle("/events/", events.InitRouter(h))
	mux.Handle("/answers/", answers.InitRouter(h))

	return http.StripPrefix("/v1", mux)
}
