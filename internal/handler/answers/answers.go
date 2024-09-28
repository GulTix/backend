package answers

import (
	"backend/internal/service/answers"
	"backend/pkg/response"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateAnswers godoc
// @Summary Create Answers
// @Tags Forms
// @Accept json
// @Produce json
// @Param body body answers.CreateBody true "Create Body"
// @Success 201 {object} answers.CreateResponse
// @Router /v1/forms/answers/ [post]
// @Security ApiKeyAuth
func (h *HandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var (
		body answers.CreateBody
		err  error
	)

	raw, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		response.ReturnInternalServerError(w, err)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(raw, &body)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	res, err := h.answerService.Create(r.Context(), body)

	if err != nil {
		if err.Error() == "duplicated data" {
			response.SetError(w, http.StatusBadRequest, fmt.Errorf("Duplicated with same user_id and event_id is already exists."))
			return
		}

		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, http.StatusCreated, res)
	return
}
