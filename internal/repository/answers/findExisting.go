package answers

import (
	"backend/internal/entity"
	"context"
)

func (r *RepositoryImpl) FindExisting (ctx context.Context, eventId string, userId string) (*bool, error) {
	var (
		answer *entity.Answer
	)

	query := `SELECT * FROM answers WHERE event_id = $1 AND user_id = $2`

	rows, err := r.db.Query(ctx, query, eventId, userId)

	if err != nil {
		return nil, err 
	}

	defer rows.Close()

	

}