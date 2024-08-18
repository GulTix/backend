package user

import (
	"backend/internal/entity"
	"context"
)

func (r *RepositoryImpl) Create(ctx context.Context, data entity.User) (*entity.User, error) {
	query := `INSERT INTO users (id, email) VALUES ($1, $2)`

	_, err := r.db.Exec(ctx, query, data.Id, data.Email)

	if err != nil {
		return nil, err
	}

	return &data, err
}
