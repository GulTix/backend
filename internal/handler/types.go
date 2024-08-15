package handler

import "backend/internal/handler/auth"

type (
	Handlers struct {
		Auth auth.Handler
	}
)
