package users

import (
	"backend/internal/entity"
	"backend/pkg/database"
	"context"
)

type (
	Repository interface {
		Create(ctx context.Context, data entity.User) (*entity.User, error)
		FindUnique(ctx context.Context, value string, column string) (*entity.User, error)
	}
	RepositoryImpl struct {
		db database.DB
	}
)
