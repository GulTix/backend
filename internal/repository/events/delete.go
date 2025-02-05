package events

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) Delete(ctx context.Context, eventId string) (*entity.Event, error) {
	query := "DELETE FROM events WHERE id = $1 RETURNING *"

	rows, err := r.db.Query(ctx, query, eventId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	event, err := pgx.CollectOneRow[entity.Event](rows, pgx.RowToStructByName[entity.Event])

	if err != nil {
		return nil, err
	}

	return &event, nil
}
