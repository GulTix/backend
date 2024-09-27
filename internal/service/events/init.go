package events

import "backend/internal/repository/events"

func NewService(eventRepo events.Repository) Service {
	return &serviceImpl{
		repo: eventRepo,
	}
}
