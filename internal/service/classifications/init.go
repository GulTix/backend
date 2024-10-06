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
		classRepo:      classificationRepo,
		allowClassRepo: allowedClassificationRepo,
	}
}
