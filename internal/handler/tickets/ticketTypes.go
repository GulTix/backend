package tickets

import (
	tickettypes "backend/internal/service/tickets"
	"backend/pkg/response"
	"encoding/json"
	"io"
	"net/http"
)

// CreateTicketType godoc
// @Summary Create Ticket Type
// @Description Create a Single Ticket Type
// @Tags Tickets
// @Accept json
// @Produce json
// @Param body body tickettypes.CreateBody true "Create Body"
// @Success 201 {object} tickettypes.CreateResponse
// @Router /v1/tickets/ [post]
// @Security ApiKeyAuth
func (h *HandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var body tickettypes.CreateBody

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

	res, err := h.ticketTypeService.CreateResponse(r.Context(), body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}
