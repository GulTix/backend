package volunteers

import (
	"backend/internal/entity"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) FindUnique(ctx context.Context, value string, column string) (*entity.Volunteer, error) {
	var volunteer entity.Volunteer

	query := fmt.Sprintf("SELECT * FROM volunteers WHERE %s = $1", column)

	rows, err := r.db.Query(ctx, query, value)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	volunteer, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Volunteer])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &volunteer, nil
}
