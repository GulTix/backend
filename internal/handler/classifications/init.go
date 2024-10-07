package classifications

import "backend/internal/service/classifications"

func NewHandler(
	classService classifications.Service,
) Handler {
	return &HandlerImpl{
		classService: classService,
	}
}
