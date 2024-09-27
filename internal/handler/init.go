package handler

import (
	"backend/internal/handler/answers"
	"backend/internal/handler/auth"
	"backend/internal/handler/events"
	"backend/internal/handler/forms"
)

func NewHandler(
	authHandler auth.Handler,
	formHandler forms.Handler,
	eventHandler events.Handler,
	answerHandler answers.Handler,
) *Handlers {
	return &Handlers{
		Auth:   authHandler,
		Form:   formHandler,
		Event:  eventHandler,
		Answer: answerHandler,
	}
}
