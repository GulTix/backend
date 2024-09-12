package handler

import (
	"backend/internal/handler/auth"
	"backend/internal/handler/forms"
)

type (
	Handlers struct {
		Auth auth.Handler
		Form forms.Handler
	}
)
