package allowedclassifications

import (
	"backend/pkg/database"
	"context"
)

type (
	Repository     interface{
		 Create(ctx context.Context, ticketId string, classificationId string) error
	}
	RepositoryImpl struct {
		db database.DB
	}
)
