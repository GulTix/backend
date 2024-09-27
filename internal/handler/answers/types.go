package answers

import (
	"backend/internal/service/answers"
	"net/http"
)

type (
	HandlerImpl struct {
		answerService answers.Service
	}

	Handler interface {
		Create(w http.ResponseWriter, r *http.Request)
	}
)
