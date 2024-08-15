package handler

import "backend/internal/handler/auth"

func NewHandler(
	authHandler auth.Handler,
) *Handlers {
	return &Handlers{
		Auth: authHandler,
	}
}
