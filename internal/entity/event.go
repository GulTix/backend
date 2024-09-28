package entity

type (
	Event struct {
		Id             string         `json:"id" db:"id"`
		Name           string         `json:"name" db:"name"`
		GoogleFormLink string         `json:"google_form_link" db:"google_form_link"`
		BlasterToken   map[string]any `json:"blaster_token" db:"blaster_token"`
		Deleted        bool           `json:"deleted" db:"deleted"`
	}
)
