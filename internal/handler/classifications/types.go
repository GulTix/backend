package classifications

import "backend/internal/service/classifications"

type (
	Handler     interface{}
	HandlerImpl struct {
		classService classifications.Service
	}
)
