package auth

import (
	"net/http"
)

// Google Login URL godoc
// @Summary Redirect to Login Goggle
// @Description Redirect URL
// @Tags Auth
// @Accept json
// @Produce json
// @Router /v1/auth/google/login [get]
func (h *HandlerImpl) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := h.oauthService.GetGoogleLoginURL()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

// Google Login URL godoc
// @Summary Redirect to Login Goggle
// @Description Redirect URL
// @Tags Auth
// @Accept json
// @Produce json
// @Router /v1 [get]
func (h *HandlerImpl) ReturnHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello World"}`))
}
