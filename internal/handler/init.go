package handler

import (
	"backend/internal/handler/auth"
	"backend/internal/handler/forms"
)

func NewHandler(
	authHandler auth.Handler,
	formHandler forms.Handler,
) *Handlers {
	return &Handlers{
		Auth: authHandler,
		Form: formHandler,
	}
}
