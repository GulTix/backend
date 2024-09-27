package volunteers

import (
	"backend/pkg/database"
)

type (
	Repository interface {
		// CreateDevVolunteer(ctx context.Context) error
	}
	RepositoryImpl struct {
		db database.DB
	}
)
