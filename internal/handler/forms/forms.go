package forms

import (
	"backend/pkg/response"
	"net/http"
)

// Get Forms
// @Summary Get Forms
// @Description Get Forms
// @Tags Forms
// @Accept json
// @Produce json
// @Router /v1/forms/ [get]
// @Security ApiKeyAuth
func (h *handlerImpl) GetForms(w http.ResponseWriter, r *http.Request) {
	// get form
	response.SetResponse(w, 200, nil, "Ini data Formnya ya bang", true)
}
