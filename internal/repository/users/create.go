package users

import (
	"backend/internal/entity"
	"context"
)

func (r *RepositoryImpl) Create(ctx context.Context, data entity.User) (*entity.User, error) {
	query := `
		INSERT INTO users (id, first_name, last_name, phone_number, email, gender) 
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.Exec(
		ctx,
		query,
		data.Id,
		data.FirstName,
		data.LastName,
		data.PhoneNumber,
		data.Email,
		data.Gender,
	)

	if err != nil {
		return nil, err
	}

	return &data, err
}
