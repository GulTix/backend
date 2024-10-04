package validations

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) Create(ctx context.Context, data entity.Validation) (*entity.Validation, error) {
	query := `INSERT INTO validations (id, answer_id, user_id, volunteer_id, classification)
	VALUES ($1, $2, $3, $4, $5) RETURNING *`

	rows, err := r.db.Query(
		ctx,
		query,
		data.Id,
		data.AnswerId,
		data.UserId,
		data.VolunteerId,
		data.Classification,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	validation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Validation])

	return &validation, nil
}
