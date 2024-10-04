package validations

import "backend/internal/service/validations"

func NewHandler(validationService validations.Service) Handler {
	return &HandlerImpl{
		validationService: validationService,
	}
}
