package events

import (
	"backend/internal/repository/events"
	"backend/pkg/oauth"
)

func NewService(eventRepo events.Repository, oauth oauth.OAuth) Service {
	return &serviceImpl{
		repo:  eventRepo,
		oauth: oauth,
	}
}
