package payments

import (
	"backend/internal/service/payments"
	"net/http"
)

type (
	Handler interface {
		PaymentCallback(w http.ResponseWriter, r *http.Request)
		TestPayment(w http.ResponseWriter, r *http.Request)
		SentMail(w http.ResponseWriter, r *http.Request)
	}

	HandlerImpl struct {
		paymentService payments.Service
	}
)
