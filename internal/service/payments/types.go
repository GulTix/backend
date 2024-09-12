package payments

import "backend/internal/entity"

type (
	Service interface {
		HandleMitransCallback(body entity.PaymentBodyCallback)
	}

	ServiceImpl struct {
	}
)
