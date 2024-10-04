package entity

import "time"

type (
	Validation struct {
		Id               string     `json:"id" db:"id"`
		AnswerId         string     `json:"answerId" db:"answer_id"`
		UserId           string     `json:"userId" db:"user_id"`
		VolunteerId      string     `json:"volunteerId" db:"volunteer_id"`
		Deleted          bool       `json:"deleted" db:"deleted"`
		PaymentEmailSent *time.Time `json:"paymentEmailSent" db:"payment_email_sent"`
		Classification   string     `json:"classification" db:"classification"`
	}
)
