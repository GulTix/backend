package answers

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) FindExisting(
	ctx context.Context,
	eventId string,
	userId string,
) (*entity.Answer, error) {
	var (
		answer entity.Answer
	)

	query := `SELECT * FROM answers WHERE event_id = $1 AND user_id = $2`

	rows, err := r.db.Query(ctx, query, eventId, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	answer, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.Answer])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &answer, nil
}
