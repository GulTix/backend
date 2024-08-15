package user

import (
	"backend/internal/entity"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) FindUnique(ctx context.Context, value string, column string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT * FROM users WHERE %s = $1", column)

	rows, err := r.db.Query(ctx, query, value)

	if err != nil {
		return user, err
	}

	defer rows.Close()

	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])

	if err != nil {
		return user, err
	}

	return user, nil
}
