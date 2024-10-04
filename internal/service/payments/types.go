package payments

import (
	"backend/internal/repository/payments"
	"backend/pkg/midtrans"
	"context"
)

type (
	Service interface {
		HandleMidtransCallback(ctx context.Context, orderId string, acquirer string) ([]byte, error)
		// SimulatePaymentAndSentEmail(body)
	}

	ServiceImpl struct {
		midtrans    midtrans.MidtransPkg
		paymentRepo payments.Repository
	}
)
