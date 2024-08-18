package auth

import (
	"backend/internal/service/oauth"
	"net/http"
)

type (
	HandlerImpl struct {
		oauthService oauth.Service
	}

	Handler interface {
		GoogleLogin(w http.ResponseWriter, r *http.Request)
		ReturnHelloWorld(w http.ResponseWriter, r *http.Request)
		GoogleCallback(w http.ResponseWriter, r *http.Request)
	}
)
