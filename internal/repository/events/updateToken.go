package events

import (
	"context"
	"log"
)

func (r *RepositoryImpl) UpdateToken(ctx context.Context, token []byte, eventId string) error {

	query := `UPDATE events SET blaster_token = $1 WHERE id = $2`

	_, err := r.db.Exec(
		ctx,
		query,
		token,
		eventId,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
