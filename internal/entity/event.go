package entity

import "time"

type (
	Event struct {
		Id             string         `json:"id" db:"id"`
		Name           string         `json:"name" db:"name"`
		GoogleFormLink string         `json:"google_form_link" db:"google_form_link"`
		BlasterToken   map[string]any `json:"blaster_token" db:"blaster_token"`
		Deleted        bool           `json:"deleted" db:"deleted"`
	}

	GoogleToken struct {
		Expiry       time.Time `json:"expiry" db:"expiry"`
		TokenType    string    `json:"token_type" db:"token_type"`
		AccessToken  string    `json:"access_token db:"access_token"`
		RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	}
)
