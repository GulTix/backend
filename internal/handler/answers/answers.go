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
// @Router /v1/forms/answers [post]
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

//func (h *HandlerImpl) Creates(w http.ReponseWriter, r *http.Request) {
//	var body answers.CreateBodyBulk
//}

// Get All Answers godoc
// @Summary Get All Answers
// @Tags Answers
// @Accept json
// @Produce json
// @Param event_id query string false "Event ID"cd
// @Router /v1/answers/ [get]
// @Security ApiKeyAuth
// @Success 200 {object} answers.FindAllResponse
// @Success 404 {object} answers.FindAllResponse
func (h *HandlerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		res *answers.FindAllResponse
		err error
	)
	eventId := r.URL.Query().Get("event_id")

	if eventId != "" {
		res, err = h.answerService.FindAllByEventId(r.Context(), eventId)

		if err != nil {
			response.ReturnInternalServerError(w, err)
			return
		}

		response.SetRawResponse(w, res.StatusCode, res)
		return
	}

	res, err = h.answerService.FindAll(r.Context())

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, res.StatusCode, res)
}
