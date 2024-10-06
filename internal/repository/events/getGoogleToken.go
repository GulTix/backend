package events

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v5"
	"golang.org/x/oauth2"
)

type EventToken struct {
	BlasterToken []byte `json:"blaster_token" db:"blaster_token"`
}

func (r *RepositoryImpl) GetBlasterToken(ctx context.Context, eventId string) (*oauth2.Token, error) {
	query := `SELECT blaster_token FROM events WHERE id = $1`

	rows, err := r.db.Query(ctx, query, eventId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	event, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[EventToken])

	if err != nil {
		return nil, err
	}

	if event.BlasterToken == nil {
		log.Println("Token is Null")
		return nil, nil
	}

	googleToken := oauth2.Token{}

	err = json.Unmarshal(event.BlasterToken, &googleToken)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &googleToken, nil
}
