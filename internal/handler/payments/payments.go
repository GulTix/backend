package payments

import (
	"backend/pkg/midtrans"
	"backend/pkg/response"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func (h *HandlerImpl) PaymentCallback(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	orderId, exists := notificationPayload["order_id"].(string)

	if !exists {
		response.SetResponse(w, http.StatusNotFound, nil, "order_id Not Found", false)
		return
	}

	acquirer, exist := notificationPayload["acquirer"].(string)

	if !exist {
		acquirer = ""
	}

	log.Println(acquirer)

	_, err = h.paymentService.HandleMidtransCallback(r.Context(), orderId, acquirer)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, 200, "ok")
}

func (h *HandlerImpl) TestPayment(w http.ResponseWriter, r *http.Request) {

	// var c = coreapi.Client{}

	// c.New(
	// 	os.Getenv("MIDTRANS_SERVER_KEY"),
	// 	midtrans.Sandbox,
	// )

	// req := &coreapi.ChargeReq{
	// 	PaymentType: coreapi.PaymentTypeQris,
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  "Order-001",
	// 		GrossAmt: 50000,
	// 	},
	// }

	// res, err := c.ChargeTransaction(req)
	// if err != nil {
	// 	response.ReturnInternalServerError(w, err)
	// 	return
	// }

	// response.SetRawResponse(w, 200, res)

	midtrans := midtrans.NewMidtrans(
		os.Getenv("MIDTRANS_SERVER_KEY"),
		os.Getenv("ENV"),
		"https://b0fd-182-253-231-201.ngrok-free.app/v1/payments/callback/",
	)

	log.Println("Midtrans Created")

	qris, err := midtrans.GenerateQRIS("Order-010", 50000)

	if err != nil {
		response.ReturnInternalServerError(w, err)
		return
	}

	response.SetRawResponse(w, http.StatusOK, qris)
}
