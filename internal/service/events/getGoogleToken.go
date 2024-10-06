package events

import (
	"context"

	"golang.org/x/oauth2"
)

func (s *serviceImpl) GetBlasterToken(ctx context.Context, eventId string) (*oauth2.Token, error) {
	return s.repo.GetBlasterToken(ctx, eventId)
}
