package handler

import (
	"backend/internal/handler/answers"
	"backend/internal/handler/auth"
	"backend/internal/handler/classifications"
	"backend/internal/handler/events"
	"backend/internal/handler/forms"
	"backend/internal/handler/payments"
	ticketTypes "backend/internal/handler/tickets"
	"backend/internal/handler/validations"
)

func NewHandler(
	authHandler auth.Handler,
	formHandler forms.Handler,
	eventHandler events.Handler,
	answerHandler answers.Handler,
	ticketTypeHandler ticketTypes.Handler,
	validationHandler validations.Handler,
	paymentHandler payments.Handler,
	classHandler classifications.Handler,
) *Handlers {
	return &Handlers{
		Auth:       authHandler,
		Form:       formHandler,
		Event:      eventHandler,
		Answer:     answerHandler,
		Ticket:     ticketTypeHandler,
		Validation: validationHandler,
		Payment:    paymentHandler,
		Class:      classHandler,
	}
}
