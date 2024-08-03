package entity

type (
	Question struct {
		Id           string `json:"id" db:"id"`
		Content      string `json:"content" db:"content"`
		QuestionType string `json:"question_type" db:"question_type"`
		FormId       string `json:"form_id" db:"form_id"`
	}
)
