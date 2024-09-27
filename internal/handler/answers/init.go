package answers

import "backend/internal/service/answers"

func NewHandler(answerService answers.Service) Handler {
	return &HandlerImpl{
		answerService: answerService,
	}
}
