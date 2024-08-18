package oauth

import (
	"context"
	"io"
	"log"
	"net/url"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewClient() OAuth {
	client := oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	return &oauthImpl{
		client: &client,
	}
}

func (o *oauthImpl) GetRedirectURL() string {
	URL, err := url.Parse(o.client.Endpoint.AuthURL)
	if err != nil {
		log.Println("Parse: " + err.Error())
	}

	parameters := url.Values{}
	parameters.Add("client_id", o.client.ClientID)
	parameters.Add("scope", strings.Join(o.client.Scopes, " "))
	parameters.Add("redirect_uri", o.client.RedirectURL)
	parameters.Add("response_type", "code")
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	return url
}

func (o *oauthImpl) GetToken(ctx context.Context,code string) (*oauth2.Token, error) {
	return o.client.Exchange(ctx, code)
}

func (o *oauthImpl) GetInfo(ctx context.Context,token *oauth2.Token, url string) ([]byte, error) {
	client := o.client.Client(ctx, token)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err 
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}