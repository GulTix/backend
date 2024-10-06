package entity

type (
	Event struct {
		Id             string `json:"id" db:"id"`
		Name           string `json:"name" db:"name"`
		GoogleFormLink string `json:"google_form_link" db:"google_form_link"`
		Deleted        bool   `json:"deleted" db:"deleted"`
	}
)
