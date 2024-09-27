package events

import (
	"backend/internal/service/events"
	"net/http"
)

type (
	HandlerImpl struct {
		eventService events.Service
	}

	Handler interface {
		Create(w http.ResponseWriter, r *http.Request)
		FindAll(w http.ResponseWriter, r *http.Request)
	}
)
