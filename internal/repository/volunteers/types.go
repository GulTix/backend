package volunteers

import (
	"backend/internal/entity"
	"backend/pkg/database"
	"context"
)

type (
	Repository interface {
		CreateDevVolunteer(ctx context.Context) error
		FindUnique(ctx context.Context, value string, column string) (*entity.Volunteer, error)
	}
	RepositoryImpl struct {
		db database.DB
	}
)
