package answers

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) FindAll(ctx context.Context) ([]entity.Answer, error) {
	var (
		answers []entity.Answer
	)

	query := "SELECT * FROM answers"

	rows, err := r.db.Query(ctx, query)

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
