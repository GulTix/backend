package events

import (
	"backend/internal/service/classifications"
	"backend/pkg/response"
	"encoding/json"
	"io"
	"net/http"
)

// CreateClassification godoc
// @Summary Create a classification
// @Description Create a classification
// @Tags Events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Param body body classifications.CreateBody true "Create Classification"
// @Router /v1/events/{id}/classifications/ [post]
// @Security ApiKeyAuth
func (h *HandlerImpl) CreateClassification(w http.ResponseWriter, r *http.Request) {
	var body classifications.CreateBody

	eventId := r.PathValue("id")

	raw, err := io.ReadAll(r.Body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	err = json.Unmarshal(raw, &body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	if body.EventId == "" {
		body.EventId = eventId
	}

	resp, err := h.classService.Create(r.Context(), body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, resp.StatusCode, resp)
}
