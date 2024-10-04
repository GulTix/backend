package tickets

import (
	tickettypes "backend/internal/service/tickets"
	"net/http"
)

type (
	Handler interface {
		Create(w http.ResponseWriter, r *http.Request)
	}

	HandlerImpl struct {
		ticketTypeService tickettypes.Service
	}
)
