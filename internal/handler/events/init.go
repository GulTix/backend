package events

import (
	"backend/internal/service/classifications"
	"backend/internal/service/events"
	"backend/internal/service/tickets"
)

func NewHandler(
	eventService events.Service,
	ticketService tickets.Service,
	classService classifications.Service,
) Handler {
	return &HandlerImpl{
		eventService:  eventService,
		ticketService: ticketService,
		classService:  classService,
	}
}
