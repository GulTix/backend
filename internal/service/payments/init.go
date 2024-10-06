package payments

import (
	"backend/internal/repository/payments"
	"backend/internal/service/events"
	"backend/pkg/midtrans"
)

func NewService(
	midtrans midtrans.MidtransPkg,
	paymentRepo payments.Repository,
	eventService events.Service,
) Service {
	return &ServiceImpl{
		midtrans:     midtrans,
		paymentRepo:  paymentRepo,
		eventService: eventService,
	}
}
