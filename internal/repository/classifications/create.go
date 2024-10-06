package classifications

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) Create(ctx context.Context, data entity.Classification) (*entity.Classification, error) {
	query := `INSERT INTO classifications (id, event_id, name)
	VALUES ($1, $2, $3) RETURNING *`

	rows, err := r.db.Query(
		ctx,
		query,
		data.Id,
		data.EventId,
		data.Name,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	classification, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Classification])

	if err != nil {
		return nil, err
	}

	return &classification, nil
}
