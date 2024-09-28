package events

import (
	"backend/internal/service/events"
	"backend/pkg/response"
	"encoding/json"
	"io"
	"net/http"
)

// CreateEvent godoc

// @Summary Create Event
// @Description Create a Single Event
// @Tags Events
// @Accept json
// @Produce json
// @Param body body events.CreateBody true "Create Body"
// @Success 201 {object} events.CreateResponse
// @Router /v1/events/ [post]
// @Security ApiKeyAuth
func (h *HandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var (
		body events.CreateBody
	)

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

	res, err := h.eventService.Create(r.Context(), body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, http.StatusCreated, res)
}

// Get All Events godoc
// @Summary Get All Events
// @Description Get All Event
// @Tags Events
// @Accept json
// @Produce json
// @Router /v1/events/ [get]
// @Security ApiKeyAuth
func (h *HandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.eventService.FindAll(r.Context())

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, http.StatusOK, res)
}
