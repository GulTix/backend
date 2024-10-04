package events

import (
	"backend/pkg/response"
	"net/http"
)

// Get All Ticket Types godoc
// @Summary Get All Ticket Types
// @Description Get All Ticket Types
// @Tags Events
// @Accept json
// @Produce json
// @Router /v1/events/{id}/tickets [get]
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
