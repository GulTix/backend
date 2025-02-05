package classifications

import (
	"backend/internal/service/classifications"
	"backend/pkg/response"
	"encoding/json"
	"io"
	"net/http"
)

// CreateClassification godoc
// @Summary Create a classification
// @Description Creat a classification
// @Tags Classifications
// @Accapet json
// @Produce json
// @Param body body classifications.CreateBody true "Create Classification"
// @Router /v1/classifications [post]
// @Security ApiKeyAuth
// @Success 200 {object} classifications.CreateResponse
func (h *HandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var body classifications.CreateBody

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

	resp, err := h.classService.Create(r.Context(), body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, resp.StatusCode, resp)
}

// FindAll godoc
// @Summary Find all classifications
// @Description Find all classifications
// @Tags Classifications
// @Produce json
// @Param eventId query string true "Event ID"
// @Router /v1/classifications [get]
// @Security ApiKeyAuth
// @Success 200 {object} classifications.FindAllResponse
func (h *HandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	eventId := r.URL.Query().Get("eventId")

	resp, err := h.classService.FindAll(r.Context(), eventId)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, resp.StatusCode, resp)
}
