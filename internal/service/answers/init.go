package answers

import (
	"backend/internal/repository/answers"
	"backend/internal/repository/users"
)

func NewService(answerRepo answers.Repository, userRepo users.Repository) Service {
	return &serviceImpl{
		repo:     answerRepo,
		userRepo: userRepo,
	}
}
