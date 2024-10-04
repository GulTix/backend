package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type (
	MidtransPkg interface {
		GenerateQRIS(orderId string, amount int64) (*coreapi.ChargeResponse, error)
		CheckTransactionStatus(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error)
	}

	midtransImpl struct {
		client coreapi.Client
	}
)
