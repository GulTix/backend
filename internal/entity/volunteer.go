package entity

type (
	Volunteer struct {
		Id string `json:"id" db:"id"`
		Username string `json:"username" db:"username"`
		Email string `json:"email" db:"email"`
		Role string `json:"role" db:"role"`
		Deleted bool `json:"deleted" db:"deleted"`
	}
)