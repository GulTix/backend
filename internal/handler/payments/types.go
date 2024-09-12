package payments

import (
	"backend/internal/service/payments"
	"net/http"
)

type (
	Handler interface {
		PaymentCallback(w http.ResponseWriter, r *http.Request)
	}

	HandlerImpl struct {
		paymentService payments.Service
	}
)
