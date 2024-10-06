package midtrans

import (
	"backend/internal/entity"
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

func (m *midtransImpl) GenerateQRIS(orderId string, amount int64, userData entity.User, expired int, items *[]midtrans.ItemDetails) (*coreapi.ChargeResponse, error) {
	reqRaw := coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeQris,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: amount,
		},
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: expired,
			Unit:           "hour",
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: userData.FirstName,
			LName: userData.LastName,
			Email: userData.Email,
		},
		Items: items,
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
