package auth

import (
	"backend/pkg/response"
	"fmt"
	"net/http"
)

// DebugLogin godoc
// @Summary Debug Login
// @Description Debug Login
// @Param email query string true "Email"
// @Tags Auth
// @Accept json
// @Produce json
// @Router /v1/auth/debug/login [post]
func (h *HandlerImpl) DevLogin(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if email == "" {
		response.SetError(w, http.StatusBadRequest, fmt.Errorf("No Email Provided"))
		return
	}

	res, err := h.authService.ReturnDevToken(r.Context(), email)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}
