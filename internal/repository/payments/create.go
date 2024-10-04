package payments

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) Create(ctx context.Context, data entity.Payment) (*entity.Payment, error) {
	query := `INSERT INTO payments (id, ticket_type_id, answer_id, qris_url, status)
	VALUES ($1, $2, $3, $4, $5) RETURNING *`

	rows, err := r.db.Query(
		ctx,
		query,
		data.Id,
		data.TicketTypeId,
		data.AnswerId,
		data.QrisUrl,
		data.Status,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	payment, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Payment])

	return &payment, nil
}
