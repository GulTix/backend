package answers

import (
	"backend/internal/entity"
	"context"
)

func (r *RepositoryImpl) Create(
	ctx context.Context,
	data entity.Answer,
) (*entity.Answer, error) {
	query := `INSERT INTO answers (id, event_id, user_id, raw_data)
	VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(
		ctx,
		query,
		data.Id,
		data.EventId,
		data.UserId,
		data.RawData,
	)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
