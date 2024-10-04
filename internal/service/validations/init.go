package validations

import "backend/internal/repository/validations"

func NewService(validationRepo validations.Repository) Service {
	return &serviceImpl{
		repo: validationRepo,
	}
}
