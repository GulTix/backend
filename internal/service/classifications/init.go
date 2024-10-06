package classifications

import (
	"backend/internal/repository/allowedclassifications"
	"backend/internal/repository/classifications"
)

func NewService(
	classificationRepo classifications.Repository,
	allowedClassificationRepo allowedclassifications.Repository,
) Service {
	return &ServiceImpl{
		classificationRepo:        classificationRepo,
		allowedClassificationRepo: allowedClassificationRepo,
	}
}
