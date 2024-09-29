package answers

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) FindAllByEventId(ctx context.Context, eventId string) ([]entity.Answer, error) {
	var (
		answers []entity.Answer
	)

	query := `SELECT * FROM answers WHERE deleted = FALSE AND event_id = $1`

	rows, err := r.db.Query(ctx, query, eventId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	answers, err = pgx.CollectRows(rows, pgx.RowToStructByName[entity.Answer])

	if err != nil {
		return nil, err
	}

	return answers, nil
}
