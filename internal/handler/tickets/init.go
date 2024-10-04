package tickets

import "backend/internal/service/tickets"

func NewHandler(ticketTypeService tickets.Service) Handler {
	return &HandlerImpl{
		ticketTypeService: ticketTypeService,
	}
}
