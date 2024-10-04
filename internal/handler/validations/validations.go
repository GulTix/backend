package validations

import (
	"backend/internal/service/validations"
	"backend/pkg/response"
	"encoding/json"
	"io"
	"net/http"
)

// Create Validations godoc
// @Summary Create Validations
// @Description Create a Single Validation
// @Tags Validations, Events
// @Accept json
// @Produce json
// @Param body body validations.CreateBody true "Create Body"
// @Success 201 {object} validations.CreateResponse
// @Router /v1/validations/ [post]
// @Router /v1/events/{id}/validations/ [post]
// @Security ApiKeyAuth
func (h *HandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var body validations.CreateBody

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

	res, err := h.validationService.CreateResponse(r.Context(), body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}
