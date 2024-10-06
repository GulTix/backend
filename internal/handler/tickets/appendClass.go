package tickets

import (
	"backend/internal/service/classifications"
	"backend/pkg/response"
	"encoding/json"
	"io"
	"net/http"
)

// AppendAllowedClassification godoc
// @Summary Append allowed classification
// @Description Append allowed classification
// @Tags Tickets
// @Param body body classifications.AppendBody true "Append Allowed Classification"
// @Accept json
// @Produce json
// @Router /v1/tickets/classifications/ [post]
// @Security ApiKeyAuth
func (h *HandlerImpl) AppendAllowedClassification(w http.ResponseWriter, r *http.Request) {
	var body classifications.AppendBody

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

	resp, err := h.classService.AppendAllowed(r.Context(), body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, resp.StatusCode, resp)
}
