package events

import (
	"backend/internal/entity"
	"context"
)

func (r *RepositoryImpl) Create(ctx context.Context, data entity.Event) (*entity.Event, error) {
	query := `
		INSERT INTO events (id, "name", bevy_link, google_form_link)
		VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(
		ctx,
		query,
		data.Id,
		data.Name,
		data.BevyLink,
		data.GoogleFormLink,
	)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
