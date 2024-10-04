package tickets

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) Create(ctx context.Context, data entity.TicketType) (*entity.TicketType, error) {
	query := `INSERT INTO ticket_type (id, name, price, event_id, bevy_link, quota)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`

	rows, err := r.db.Query(
		ctx,
		query,
		data.Id,
		data.Name,
		data.Price,
		data.EventId,
		data.BevyLink,
		data.Quota,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ticketType, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.TicketType])

	if err != nil {
		return nil, err
	}

	return &ticketType, nil
}
