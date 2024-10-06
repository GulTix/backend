package payments

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) UpdateQrisUrl(ctx context.Context, qrisUrl string, paymentId string) (*entity.Payment, error) {
	query := `UPDATE payments SET qris_url = $1, updated_at = NOW() WHERE id = $2 RETURNING *`

	rows, err := r.db.Query(ctx, query, qrisUrl, paymentId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	payment, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.Payment])

	if err != nil {
		return nil, err
	}

	return &payment, nil
}
