package tickets

import (
	"backend/internal/service/classifications"
	"backend/internal/service/tickets"
)

func NewHandler(
	ticketTypeService tickets.Service,
	classService classifications.Service,
) Handler {
	return &HandlerImpl{
		ticketTypeService: ticketTypeService,
		classService:      classService,
	}
}
