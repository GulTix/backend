package v1

import (
	"backend/internal/handler"
	"backend/internal/router/v1/answers"
	"backend/internal/router/v1/auth"
	"backend/internal/router/v1/events"
	"backend/internal/router/v1/forms"
	"backend/internal/router/v1/payments"
	"backend/internal/router/v1/tickets"
	"backend/internal/router/v1/validations"
	"net/http"
)

func InitV1Router(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Auth.ReturnHelloWorld)

	mux.HandleFunc("POST /events", h.Event.Create)
	mux.HandleFunc("GET /events", h.Event.FindAll)

	mux.HandleFunc("GET /tickets", h.Ticket.FindAll)
	mux.HandleFunc("POST /tickets", h.Ticket.Create)

	mux.Handle("/auth/", auth.InitRouter(h))
	mux.Handle("/forms/", forms.InitRouter(h))
	mux.Handle("/events/", events.InitRouter(h))
	mux.Handle("/answers/", answers.InitRouter(h))
	mux.Handle("/validations/", validations.InitRouter(h))
	mux.Handle("/tickets/", tickets.InitRouter(h))
	mux.Handle("/payments/", payments.InitRouter(h))

	return http.StripPrefix("/v1", mux)
}
