package payments

import (
	"backend/internal/repository/payments"
	"backend/internal/service/events"
	"backend/pkg/midtrans"
	"backend/pkg/oauth"
	"context"

	"golang.org/x/oauth2"
)

type (
	Service interface {
		HandleMidtransCallback(ctx context.Context, orderId string, payload map[string]any) ([]byte, error)
		TestParsingToken(ctx context.Context, eventId string) (*oauth2.Token, error)
		TestGenerateQRIS(ctx context.Context) (any, error)
		// SimulatePaymentAndSentEmail(body)
	}

	ServiceImpl struct {
		midtrans     midtrans.MidtransPkg
		paymentRepo  payments.Repository
		eventService events.Service
		oauthGmail   oauth.OAuth
	}
)
