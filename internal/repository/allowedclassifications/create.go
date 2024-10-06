package allowedclassifications

import (
	"context"

	"github.com/google/uuid"
)

func (r *RepositoryImpl) Create(ctx context.Context, ticketId string, classificationId string) error {
	query := "INSERT INTO allowed_classifications (id, ticket_id, classification_id) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(ctx, query, uuid.NewString(), ticketId, classificationId)

	if err != nil {
		return err
	}

	return nil
}
