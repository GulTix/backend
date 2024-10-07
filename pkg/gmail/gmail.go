package gmail

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"

	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type (
	Gmail interface{
		SendMailWithHTMLTemplate(ctx context.Context, param GmailServiceBody, templatePath string) (bool, error)
	}

	gmailImpl struct {
		client *gmail.Service
	}

	GmailServiceBody struct {
		To           string `json:"to"`
		SenderName   string
		SenderEmail  string
		Message      *string `json:"message"`
		SubjectData  string  `json:"subjectData"`
		TemplateData map[string]any
	}
)

func NewClient(ctx context.Context, tokenSource oauth2.TokenSource) Gmail {
	service, err := gmail.NewService(ctx, option.WithTokenSource(tokenSource))

	if err != nil {
		log.Println("Errer Creating Gmail Service")
		log.Print(err)
	}

	return &gmailImpl{
		client: service,
	}
}

func (g *gmailImpl) SendMailWithHTMLTemplate(ctx context.Context, param GmailServiceBody, templatePath string) (bool, error) {
	// Create Header
	subject := "Subject: " + param.SubjectData + "\n"
	from := "From: " + param.SenderEmail + " <" + param.SenderEmail + ">\n"
	to := "To: " + param.To + "\n"
	contentType := "Content-Type: text/html; charset=\"UTF-8\"\n"

	// Generating Template
	templ, err := template.ParseFiles(templatePath)

	if err != nil {
		return false, fmt.Errorf("Error loading template")
	}

	var body bytes.Buffer

	err = templ.Execute(&body, param.TemplateData)

	if err != nil {
		return false, fmt.Errorf("Error parsing data to template")
	}

	msgString := from + to + subject + contentType + "\n" + body.String()

	msg := []byte(msgString)

	message := gmail.Message{
		Raw: base64.URLEncoding.EncodeToString(msg),
	}

	_, err = g.client.Users.Messages.Send("me", &message).Do()

	if err != nil {
		return false, fmt.Errorf("Error sending Email")
	}

	return true, nil
}
