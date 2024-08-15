package oauth

import (
	"backend/pkg/oauth"
)

type (
	Service interface {
		GetGoogleLoginURL() string
		// GoogleLogin(ctx context.Context) error
		// GoogleCallback(ctx context.Context) error
	}

	ServiceImpl struct {
		oauth oauth.OAuth
	}
)
