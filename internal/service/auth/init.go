package auth

import "backend/internal/repository/volunteers"

func NewService(volunteerRepo volunteers.Repository) Service {
	return &serviceImpl{
		volunteerRepo: volunteerRepo,
	}
}
