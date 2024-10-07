package payments

import (
	"backend/internal/entity"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"golang.org/x/oauth2"
)

func (s *ServiceImpl) TestParsingToken(ctx context.Context, eventId string) (*oauth2.Token, error) {
	// token, err := s.eventService.GetBlasterToken(ctx, eventId)

	// if err != nil {
	// 	return nil, err
	// }

	// tokenSource := s.oauthGmail.GetTokenSource(ctx, token)

	// gmailClient := gmail.NewClient(ctx, tokenSource)

	// gmailClient.SendMailWithHTMLTemplate(ctx, gmail.GmailServiceBody{
	// 	To: "abdulaziz@bangunindo.com",
	// 	SenderName : "Google Developer Group Cloud Jakarta",
	// 	SenderEmail: "muhammadabdulazizalghofari",
	// 	SubjectData : "Confirmation Payment Email",
	// 	TempaletData : map[string] {
	// 		"Name": "Razan",
	// 		"EventName": "DevFest Cloud Jakarta 2024",
	// 		"TicketPrice": 50000,
	// 		"PaymentDeadline": 48,
	// 		"QRISUrl": "https://api.midtrans",

	// 	}
	// })

	return s.eventService.GetBlasterToken(ctx, eventId)
}

func (s *ServiceImpl) TestGenerateQRIS(ctx context.Context) (any, error) {
	payment := entity.Payment{
		Id:           uuid.NewString(),
		AnswerId:     "4b892441-2d6a-467f-aa45-5fcb0c6658d5",
		TicketTypeId: "571061fd-8874-4bba-97b8-d075a16fed34",
	}

	paymentReturned, err := s.paymentRepo.Create(ctx, payment)

	if err != nil {
		return nil, err
	}

	log.Println(paymentReturned)

	res, err := s.midtrans.GenerateQRIS(payment.Id, 100000, entity.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "abdulaziz@bangunindo.com",
	},
		48,
		&[]midtrans.ItemDetails{
			{
				ID:    "123",
				Name:  "Kopi",
				Price: 25000,
				Qty:   2,
			},
			{
				ID:    "124",
				Name:  "Teh",
				Price: 25000,
				Qty:   2,
			},
		})

	qrurl := &res.Actions[0].URL
	// Update QRIS URL
	paymentReturned, err = s.paymentRepo.UpdateQrisUrl(ctx, *qrurl, payment.Id)

	if err != nil {
		return nil, err
	}

	return res, err

}
