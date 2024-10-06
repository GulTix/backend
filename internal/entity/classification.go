package entity

type (
	Classification struct {
		Id      string `json:"id" db:"id"`
		Name    string `json:"name" db:"name"`
		EventId string `json:"event_id" db:"event_id"`
	}
)
