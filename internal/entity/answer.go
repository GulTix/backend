package entity

type (
	Answer struct {
		Id         string `json:"id" db:"id"`
		EventId    string `json:"event_id" db:"event_id"`
		UserId     string `json:"user_id" db:"user_id"`
		RawData    map[string]any `json:"raw_data" db:"raw_data"`
		Deleted    bool   `json:"deleted" db:"deleted"`
	}
)
