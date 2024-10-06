package classifications

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) FindAll(ctx context.Context) ([]entity.Classification, error) {
	query := "SELECT * FROM classifications"

	rows, err := r.db.Query(ctx, query)

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
