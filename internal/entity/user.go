package entity

type (
	User struct {
		Id          string `json:"id" db:"id"`
		FirstName   string `json:"first_name" db:"first_name"`
		LastName    string `json:"last_name" db:"last_name"`
		PhoneNumber string `json:"phone_number" db:"phone_number"`
		Gender      string `json:"gender" db:"gender"`
		Email       string `json:"email" db:"email"`
		Deleted     bool   `json:"deleted" db:"deleted"`
	}
)
