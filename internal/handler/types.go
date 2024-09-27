package handler

import (
	"backend/internal/handler/answers"
	"backend/internal/handler/auth"
	"backend/internal/handler/events"
	"backend/internal/handler/forms"
)

type (
	Handlers struct {
		Auth   auth.Handler
		Form   forms.Handler
		Event  events.Handler
		Answer answers.Handler
	}
)
