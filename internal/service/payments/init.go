package payments

import (
	"backend/internal/repository/payments"
	"backend/internal/service/events"
	"backend/pkg/midtrans"
	"backend/pkg/oauth"
)

func NewService(
	midtrans midtrans.MidtransPkg,
	paymentRepo payments.Repository,
	eventService events.Service,
	oauthGmail oauth.OAuth,
) Service {
	return &ServiceImpl{
		midtrans:     midtrans,
		paymentRepo:  paymentRepo,
		eventService: eventService,
		oauthGmail:   oauthGmail,
	}
}
