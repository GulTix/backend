package classifications

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) FindAllByEventId(ctx context.Context, eventId string) ([]entity.Classification, error) {
	query := "SELECT id, name, event_id FROM classifications WHERE event_id = $1"

	rows, err := r.db.Query(ctx, query, eventId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	classifications, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Classification])

	if err != nil {
		return nil, err
	}

	return classifications, nil
}
