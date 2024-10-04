package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func (s *ServiceImpl) GetGoogleLoginURL() string {
	return s.oauth.GetRedirectURL()
}

func (s *ServiceImpl) ReturnGoogleCallbackResponse(ctx context.Context, code string) (*GoogleCallbackResponse, error) {
	token, err := s.oauth.GetToken(ctx, code)

	if err != nil {
		log.Printf("[Oauth][ReturnGoogleCallbackResponse][GetToken] Error Getting Google Token +%v\n", err)
		return nil, err
	}

	log.Printf("%+v", token)

	data, err := s.oauth.GetInfo(ctx, token, "https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		log.Printf("[Oauth][ReturnGoogleCallbackResponse][GetInfo] Error Getting Google Response Data +%v\n", err)
		return nil, err
	}

	userGoogle := GoogleUserData{}
	err = json.Unmarshal(data, &userGoogle)

	if err != nil {
		log.Printf("[Oauth][ReturnGoogleCallbackResponse][Unmarshal] Error Unmarshaling Google Response Data +%v\n", err)
		return nil, err
	}

	volunteer, err := s.volunteerRepo.FindUnique(ctx, userGoogle.Email, "email")

	if err != nil {
		log.Printf("[Oauth][ReturnGoogleCallbackResponse][FindUnique] Error Finding User +%v\n", err)
		return nil, err
	}

	if volunteer == nil {
		return nil, fmt.Errorf("User not found")
	}

	signedToken, err := s.authService.GenerateToken(*volunteer)

	if err != nil {
		log.Printf("[Oauth][ReturnGoogleCallbackResponse][GenerateToken] Error Generating Token +%v\n", err)
		return nil, err
	}

	loginData := LoginData{
		User:  *volunteer,
		Token: *signedToken,
	}

	return &GoogleCallbackResponse{
		Success: true,
		Message: "Successfully Logged In",
		Data:    loginData,
	}, nil
}
