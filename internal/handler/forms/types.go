package forms

import "net/http"

type (
	handlerImpl struct {
	}

	Handler interface {
		GetForms(w http.ResponseWriter, r *http.Request)
	}
)
