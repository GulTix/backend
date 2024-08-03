package user

import "backend/pkg/database"

type (
	Repository     interface{}
	RepositoryImpl struct {
		db database.DB
	}
)
