package events

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) FindAll(ctx context.Context) ([]entity.Event, error) {
	var (
		events []entity.Event
	)
	query := `
		SELECT id, name, google_form_link, deleted FROM events
	`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events, err = pgx.CollectRows(rows, pgx.RowToStructByName[entity.Event])

	if err != nil {
		return nil, err
	}

	return events, nil
}
