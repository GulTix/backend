package tickets

import (
	"backend/internal/service/classifications"
	tickettypes "backend/internal/service/tickets"
	"net/http"
)

type (
	Handler interface {
		Create(w http.ResponseWriter, r *http.Request)
		FindAll(w http.ResponseWriter, r *http.Request)
	}

	HandlerImpl struct {
		ticketTypeService tickettypes.Service
		classService      classifications.Service
	}
)
