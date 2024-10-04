package events

import (
	"backend/internal/service/events"
	"backend/pkg/response"
	"encoding/json"
	"io"
	"log"
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

// Callback For Google Token godoc
// @Summary Callback For Google Token
// @Description Callback For Google Token
// @Tags Events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Param code query string true "Google Code"
// @Router /v1/events/{id}/callback/ [put]
// @Security ApiKeyAuth
func (h *HandlerImpl) SetGoogleToken(w http.ResponseWriter, r *http.Request) {
	var (
		eventId string
		code    string
	)

	eventId = r.PathValue("id")
	code = r.URL.Query().Get("code")

	log.Println(eventId)
	log.Println(code)

	res, err := h.eventService.SetGoogleTokenResponse(r.Context(), eventId, code)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}

// Google Login for Token godoc
// @Summary Google Login for Token
// @Description Get Google Login URL
// @Tags Events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Router /v1/events/{id}/login/ [get]
// @Security ApiKeyAuth
func (h *HandlerImpl) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := h.eventService.GetGoogleLoginURL()
	response.SetResponse(w, 200, url, "Ini Redirect URL nya", true)
	// http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}
