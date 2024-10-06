package payments

import (
	"backend/internal/entity"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) UpdateStatus(ctx context.Context, paymentId string, status string, payload map[string]any) error {
	query := `UPDATE payments SET status = $1, updated_at = NOW(), midtrans_payload = $2 WHERE id = $3 RETURNING *`

	rows, err := r.db.Query(ctx, query, status, payload, paymentId)

	if err != nil {
		return err
	}

	defer rows.Close()

	payment, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[entity.Payment])

	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v", payment)

	return nil
}
