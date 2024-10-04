package payments

import (
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

}
