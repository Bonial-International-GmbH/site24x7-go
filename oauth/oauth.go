package oauth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	TokenURL = "https://accounts.zoho.com/oauth/v2/token"

	TokenType = "Zoho-oauthtoken"
)

type Config struct {
	*oauth2.Config

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string
}

func NewConfig(clientID, clientSecret, refreshToken string) *Config {
	return &Config{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				TokenURL: TokenURL,
			},
		},
		RefreshToken: refreshToken,
	}
}

func (c *Config) Client(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, c.TokenSource(ctx))
}

func (c *Config) TokenSource(ctx context.Context) oauth2.TokenSource {
	t := &oauth2.Token{
		RefreshToken: c.RefreshToken,
		TokenType:    TokenType,
	}

	return c.Config.TokenSource(ctx, t)
}
