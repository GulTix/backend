package payments

import (
	"context"
	"log"
)

func (s *ServiceImpl) HandleMidtransCallback(ctx context.Context, orderId string, payload map[string]any) ([]byte, error) {
	transactionStatusResp, err := s.midtrans.CheckTransactionStatus(orderId)

	if err != nil {
		log.Println("Masuk error bang")
		log.Println(err)
		return nil, err
	}

	log.Println("Gak lewat error bang")

	status := transactionStatusResp.TransactionStatus
	gulTixStatus := "pending"

	log.Println(transactionStatusResp)
	log.Println(status)

	if status == "capture" {
		if transactionStatusResp.FraudStatus == "challenge" {
			gulTixStatus = "challenge"
		}
		if transactionStatusResp.FraudStatus == "accept" {
			gulTixStatus = "success"
		}
	}

	if status == "settlement" {
		gulTixStatus = "success"
	}

	if status == "cancel" {
		gulTixStatus = "cancel"
	}

	if status == "expire" {
		gulTixStatus = "expire"
	}

	if status == "pending" {
		gulTixStatus = "pending"
	}

	log.Println("Status")
	log.Println(status)
	log.Println("GulTix Status")
	log.Println(gulTixStatus)

	e := s.paymentRepo.UpdateStatus(ctx, orderId, gulTixStatus, payload)

	if e != nil {
		log.Println("Masuk error bang")
		log.Println(e)
		return nil, e
	}

	log.Println("Paymen Update Sukses Bang")

	return []byte("ok"), nil
}
