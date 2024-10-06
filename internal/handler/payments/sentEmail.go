package payments

import (
	"backend/pkg/response"
	"net/http"
)

type (
	Params struct {
		To           string `json:"to"`
		TemplateData any    `json:"template_data"`
		SubjectData  string
		From         string
	}
)

func (h *HandlerImpl) SentMail(w http.ResponseWriter, r *http.Request) {
	eventId := r.URL.Query().Get("event_id")
	token, err := h.paymentService.TestParsingToken(r.Context(), eventId)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, http.StatusOK, token)
}
