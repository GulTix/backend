package classifications

import (
	"backend/internal/service/classifications"
	"net/http"
)

type (
	Handler interface {
		Create(w http.ResponseWriter, r *http.Request)
	}
	HandlerImpl struct {
		classService classifications.Service
	}
)
