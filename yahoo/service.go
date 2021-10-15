package yahoo

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

const clientId = "clientid"
const clientSecret = "clientsecret"
const redirectURL = "https://localhost:3000"

var (
	conf = GetOAuth2Config(clientId, clientSecret, redirectURL)
)

type YahooService interface {
	GetRedirectURL() string
	GetYahooToken(code string) *oauth2.Token
	NewYahooClient(token *oauth2.Token) *Client
}

type yahooService struct {
}

func NewService() YahooService {
	return &yahooService{}
}

func (s *yahooService) GetRedirectURL() string {
	return conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (s *yahooService) NewYahooClient(token *oauth2.Token) *Client {

	// // Use the authorization code that is pushed to the redirect
	// // URL. Exchange will do the handshake to retrieve the
	// // initial access token. The HTTP Client returned by
	// // conf.Client will refresh the token as necessary.
	// var code string
	// if _, err := fmt.Scanln(&code); err != nil {
	// 	log.Fatal(err)
	// 	return nil
	// }
	ctx := context.Background()
	client := conf.Client(ctx, token)
	return NewClient(client)
}

func (s *yahooService) GetYahooToken(code string) *oauth2.Token {
	ctx := context.Background()
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tok
}
