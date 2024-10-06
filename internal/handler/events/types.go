package events

import (
	"backend/internal/service/events"
	"backend/internal/service/tickets"
	"net/http"
)

type (
	HandlerImpl struct {
		eventService  events.Service
		ticketService tickets.Service
	}

	Handler interface {
		Create(w http.ResponseWriter, r *http.Request)
		FindAll(w http.ResponseWriter, r *http.Request)
		GoogleLogin(w http.ResponseWriter, r *http.Request)
		SetGoogleToken(w http.ResponseWriter, r *http.Request)
		FindAllTicket(w http.ResponseWriter, r *http.Request)
		CreateTicket(w http.ResponseWriter, r *http.Request)
	}
)
