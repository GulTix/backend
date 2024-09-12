package payments

import "backend/internal/service/payments"

func NewHandler(paymentService payments.Service) Handler {
	return &HandlerImpl{
		paymentService: paymentService,
	}
}
