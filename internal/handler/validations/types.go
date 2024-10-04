package validations

import "backend/internal/service/validations"

type (
	Handler interface {}
	HandlerImpl struct {
		validationService validations.Service
	}
)
