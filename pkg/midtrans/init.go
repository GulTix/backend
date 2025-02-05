package midtrans

import (
	"backend/internal/entity"
	"backend/pkg/config"
	"log"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func NewMidtrans(cfg config.Config) MidtransPkg {

	env := midtrans.Sandbox

	if cfg.Env == "prod" || cfg.Env == "production" {
		env = midtrans.Production
	}

	var client = coreapi.Client{}

	client.New(
		cfg.MidtransServerKey,
		env,
	)

	client.Options.SetPaymentOverrideNotification(cfg.BaseUrl + "/v1/payments/callback/")

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
