package events

import (
	"backend/internal/service/tickets"
	"backend/pkg/response"
	"encoding/json"
	"io"
	"net/http"
)

// Get All Ticket Types godoc
// @Summary Get All Ticket Types
// @Description Get All Ticket Types
// @Param id path string true "Event ID"
// @Tags Events
// @Accept json
// @Produce json
// @Router /v1/events/{id}/tickets/ [get]
// @Security ApiKeyAuth
func (h *HandlerImpl) FindAllTicket(w http.ResponseWriter, r *http.Request) {
	eventId := r.PathValue("id")

	res, err := h.ticketService.FindByEventIdResponse(r.Context(), eventId)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}

func (h *HandlerImpl) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var body tickets.CreateBody

	eventId := r.PathValue("id")

	raw, err := io.ReadAll(r.Body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(raw, &body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	if body.EventId == "" {
		body.EventId = eventId
	}

	res, err := h.ticketService.CreateResponse(r.Context(), body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}
