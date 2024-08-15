package oauth

import (
	"golang.org/x/oauth2"
)

type (
	OAuth interface {
		GetRedirectURL() string
	}

	oauthImpl struct {
		client *oauth2.Config
	}
)
