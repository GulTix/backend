package auth

import (
	"backend/internal/service/auth"
	"backend/internal/service/oauth"
	"net/http"
)

type (
	HandlerImpl struct {
		authService  auth.Service
		oauthService oauth.Service
	}

	Handler interface {
		GoogleLogin(w http.ResponseWriter, r *http.Request)
		ReturnHelloWorld(w http.ResponseWriter, r *http.Request)
		GoogleCallback(w http.ResponseWriter, r *http.Request)
		DevLogin(w http.ResponseWriter, r *http.Request)
	}
)
