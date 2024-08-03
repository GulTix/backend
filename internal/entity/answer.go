package entity

type (
	Answer struct {
		Id         string `json:"id" db:"id"`
		QuestionId string `json:"question_id" db:"question_id"`
		UserId     string `json:"user_id" db:"user_id"`
		FormId     string `json:"form_id" db:"form_id"`
		Value      string `json:"value" db:"value"`
	}
)
