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
