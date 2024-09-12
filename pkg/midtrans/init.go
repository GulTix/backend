package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func NewMidtrans(serverKey string, environment string, notificationUrl string) Midtrans {

	env := midtrans.Sandbox

	if environment == "production" {
		env = midtrans.Production
	}

	client := coreapi.Client{
		ServerKey: serverKey,
		Env:       env,
		Options: &midtrans.ConfigOptions{
			PaymentOverrideNotification: &notificationUrl,
			PaymentAppendNotification:   &notificationUrl,
		},
	}

	return midtransImpl{
		client: client,
	}
}

func (m *midtransImpl) GenerateQRIS(orderId string, amount int64) (*string, error) {
	req, err := m.client.ChargeTransaction(&coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeQris,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: amount,
		},
	})

	if err != nil {
		return nil, err
	}

	return &req.Actions[0].URL, nil
}
