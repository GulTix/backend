package oauth

import (
	"log"
	"net/url"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

func NewClient() OAuth {
	client := oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
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
	log.Println(URL.String())
	parameters := url.Values{}
	parameters.Add("client_id", o.client.ClientID)
	parameters.Add("scope", strings.Join(o.client.Scopes, " "))
	parameters.Add("redirect_uri", o.client.RedirectURL)
	parameters.Add("response_type", "code")
	// parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	return url
}
