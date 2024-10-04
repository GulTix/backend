package handler

import (
	"backend/internal/handler/answers"
	"backend/internal/handler/auth"
	"backend/internal/handler/events"
	"backend/internal/handler/forms"
	"backend/internal/handler/payments"
	ticketTypes "backend/internal/handler/tickets"
	"backend/internal/handler/validations"
)

type (
	Handlers struct {
		Auth       auth.Handler
		Form       forms.Handler
		Event      events.Handler
		Answer     answers.Handler
		Ticket     ticketTypes.Handler
		Validation validations.Handler
		Payment    payments.Handler
	}
)
