package entity

type (
	User struct {
		Id      string `json:"id" db:"id"`
		Email   string `json:"email" db:"email"`
		Deleted bool   `json:"deleted" db:"deleted"`
	}
)
