package forms

import (
	"backend/pkg/response"
	"net/http"
)

func (h *handlerImpl) GetForms(w http.ResponseWriter, r *http.Request) {
	// get form
	response.SetResponse(w, 200, nil, "Ini data Formnya ya bang", true)
}
