package entity

type (
	TicketType struct {
		Id                     string  `json:"id" db:"id"`
		Name                   string  `json:"name" db:"name"`
		Price                  float64 `json:"price" db:"price"`
		EventId                string  `json:"event_id" db:"event_id"`
		BevyLink               string  `json:"bevy_link" db:"bevy_link"`
		Quota                  int     `json:"quota" db:"quota"`
		PaymentExpiredDuration int     `json:"payment_expired_duration" db:"payment_expired_duration"`
		Deleted                bool    `json:"deleted" db:"deleted"`
	}
)
