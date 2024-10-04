package midtrans

import (
	"log"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func NewMidtrans(serverKey string, environment string, notificationUrl string) MidtransPkg {

	env := midtrans.Sandbox

	if environment == "prod" {
		env = midtrans.Production
	}

	var client = coreapi.Client{}

	client.New(
		serverKey,
		env,
	)

	client.Options.SetPaymentOverrideNotification(notificationUrl)

	return &midtransImpl{
		client: client,
	}
}

func (m *midtransImpl) GenerateQRIS(orderId string, amount int64) (*coreapi.ChargeResponse, error) {
	reqRaw := coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeQris,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: amount,
		},
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: 24,
			Unit:           "hour",
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: "Azei",
			LName: "Heu",
			Email: "muhammadabdulazizalghofari@gmail.com",
		},
	}

	req, err := m.client.ChargeTransaction(&reqRaw)

	if err != nil {
		log.Printf("err %+v", err)
		return nil, err
	}

	return req, nil
}

func (m *midtransImpl) CheckTransactionStatus(orderId string) (
	*coreapi.TransactionStatusResponse,
	*midtrans.Error,
) {
	transactionStatusResp, err := m.client.CheckTransaction(orderId)
	return transactionStatusResp, err
}
