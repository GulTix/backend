package oauth

import "backend/pkg/oauth"

func NewService(oauth oauth.OAuth) Service {
	return &ServiceImpl{
		oauth: oauth,
	}
}
