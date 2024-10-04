package events

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) GetGoogleToken(ctx context.Context, eventId string) (*entity.GoogleToken, error) {
	query := `SELECT blaster_token FROM events WHERE id = $1`

	rows, err := r.db.Query(ctx, query, eventId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	googleToken, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.GoogleToken])

	if err != nil {
		return nil, err
	}

	return &googleToken, nil
}
