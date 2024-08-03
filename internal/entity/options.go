package entity

type (
	ComboBoxOption struct {
		Id         string `json:"id" db:"id"`
		QuestionId string `json:"question_id" db:"question_id"`
		Option     string `json:"option" db:"option"`
	}

	RadioOption struct {
		Id         string `json:"id" db:"id"`
		QuestionId string `json:"question_id" db:"question_id"`
		Option     string `json:"option" db:"option"`
	}
)
