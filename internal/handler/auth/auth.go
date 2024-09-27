package auth

import (
	"backend/pkg/response"
	"fmt"
	"log"
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
	log.Println("Incoming")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello World"}`))
}

// Google Callbacks godoc
// @Summary Callbacks Google Login
// @Description Callbacks Google Login
// @Param code query string true "Code"
// @Tags Auth
// @Accept json
// @Produce json
// @Router /v1/auth/google/callback [get]
// @Success 200 {object} oauth.GoogleCallbackResponse
func (h *HandlerImpl) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		response.SetError(w, http.StatusBadRequest, fmt.Errorf("No Code Returned"))
		return
	}

	res, err := h.oauthService.ReturnGoogleCallbackResponse(r.Context(), code)

	if err != nil {
		response.SetError(w, http.StatusInternalServerError, err)
		return
	}

	response.SetRawResponse(w, http.StatusOK, res)
}
