package entity

import "time"

type (
	PaymentBodyCallback struct {
		TransactionType   string `json:"transaction_type"`
		TransactionTime   string `json:"transaction_time"`
		TransactionStatus string `json:"transaction_status"`
		TransactionID     string `json:"transaction_id"`
		StatusMessage     string `json:"status_message"`
		StatusCode        string `json:"status_code"`
		SignatureKey      string `json:"signature_key"`
		SettlementTime    string `json:"settlement_time"`
		PaymentType       string `json:"payment_type"`
		OrderID           string `json:"order_id"`
		MerchantID        string `json:"merchant_id"`
		Issuer            string `json:"issuer"`
		GrossAmount       string `json:"gross_amount"`
		FraudStatus       string `json:"fraud_status"`
		Currency          string `json:"currency"`
		Acquirer          string `json:"acquirer"`
	}

	Payment struct {
		Id              string         `json:"id" db:"id"`
		AnswerId        string         `json:"answer_id" db:"answer_id"`
		TicketTypeId    string         `json:"ticket_type_id" db:"ticket_type_id"`
		QrisUrl         string         `json:"qris_url" db:"qris_url"`
		Status          string         `json:"status" db:"status"`
		Deleted         bool           `json:"deleted" db:"deleted"`
		CreatedAt       time.Time      `json:"created_at" db:"created_at"`
		UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
		MidtransPayload map[string]any `json:"midtrans_payload" db:"midtrans_payload"`
	}
)
