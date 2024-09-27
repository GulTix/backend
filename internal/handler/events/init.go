package events

import "backend/internal/service/events"

func NewHandler(eventService events.Service) Handler {
	return &HandlerImpl{
		eventService: eventService,
	}
}
