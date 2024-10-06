package tickets

import (
	"backend/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) FindAllByEventId(ctx context.Context, eventId string) ([]entity.TicketType, error) {
	query := `SELECT * FROM ticket_type WHERE event_id = $1`

	rows, err := r.db.Query(ctx, query, eventId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ticketTypes, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.TicketType])

	if err != nil {
		return nil, err
	}

	return ticketTypes, nil
}
