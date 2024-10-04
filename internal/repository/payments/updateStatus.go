package payments

import (
	"context"
)

func (r *RepositoryImpl) UpdateStatus(ctx context.Context, paymentId string, status string, acquirer string) error {
	query := `UPDATE payments SET status = $1, updated_at = NOW(), acquirer = $3 WHERE id = $2 RETURNING *`

	_, err := r.db.Query(ctx, query, status, paymentId, acquirer)

	if err != nil {
		return err
	}

	return nil
}
