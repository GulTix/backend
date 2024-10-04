package payments

import (
	"backend/internal/repository/payments"
	"backend/pkg/midtrans"
)

func NewService(midtrans midtrans.MidtransPkg, paymentRepo payments.Repository) Service {
	return &ServiceImpl{
		midtrans:    midtrans,
		paymentRepo: paymentRepo,
	}
}
