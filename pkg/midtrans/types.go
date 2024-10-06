package midtrans

import (
	"backend/internal/entity"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type (
	MidtransPkg interface {
		GenerateQRIS(orderId string, amount int64, userData entity.User, expired int, items *[]midtrans.ItemDetails) (*coreapi.ChargeResponse, error)
		CheckTransactionStatus(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error)
	}

	midtransImpl struct {
		client coreapi.Client
	}
)
