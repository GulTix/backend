package events

import (
	"context"
	"encoding/json"
	"log"
)

func (s *serviceImpl) SetGoogleTokenResponse(ctx context.Context, eventId string, code string) (*SetGoogleTokenResponse, error) {
	token, err := s.oauth.GetToken(ctx, code)

	if err != nil {
		return nil, err
	}

	log.Printf("token %+v", token)

	parsedToken, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	log.Println("parsedToken", parsedToken)

	err = s.repo.UpdateToken(ctx, parsedToken, eventId)

	if err != nil {
		return nil, err
	}

	return &SetGoogleTokenResponse{
		Success:    true,
		Message:    "Token has been set",
		StatusCode: 200,
	}, nil
}

func (s *serviceImpl) GetGoogleLoginURL() string {
	return s.oauth.GetRedirectURL()
}
