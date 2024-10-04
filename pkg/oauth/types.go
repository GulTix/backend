package oauth

import (
	"context"

	"golang.org/x/oauth2"
)

type (
	OAuth interface {
		GetRedirectURL() string
		GetToken(ctx context.Context, code string) (*oauth2.Token, error)
		GetInfo(ctx context.Context, token *oauth2.Token, url string) ([]byte, error)
		GetTokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource
	}

	oauthImpl struct {
		client *oauth2.Config
	}
)
