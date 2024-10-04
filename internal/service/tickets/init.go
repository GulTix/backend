package tickets

import tickettype "backend/internal/repository/tickets"

func NewService(ticketTypeRepo tickettype.Repository) Service {
	return &serviceImpl{
		repo: ticketTypeRepo,
	}
}
