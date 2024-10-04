package events

import (
	"backend/internal/service/events"
	"backend/internal/service/tickets"
)

func NewHandler(
	eventService events.Service,
	ticketService tickets.Service,
) Handler {
	return &HandlerImpl{
		eventService:  eventService,
		ticketService: ticketService,
	}
}
