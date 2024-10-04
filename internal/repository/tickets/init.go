package tickets

import "backend/pkg/database"

func NewRepository(db database.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}
