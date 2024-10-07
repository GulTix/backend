package payments

import (
	"backend/internal/handler"
	"net/http"
)

func InitRouter(h *handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", h.Payment.TestPayment)
	mux.HandleFunc("POST /callback", h.Payment.PaymentCallback)
	mux.HandleFunc("GET /email", h.Payment.SentMail)

	return http.StripPrefix("/payments", mux)
}
