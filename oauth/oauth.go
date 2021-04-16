package oauth

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	// TokenType is the type used in the Authorization header next to the
	// access token.
	TokenType = "Zoho-oauthtoken"
)

// Config is an OAuth config that is also aware of the refresh token.
type Config struct {
	*oauth2.Config

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string
}

// NewConfig creates a new *Config for the provided client credentials.
func NewConfig(clientID, clientSecret, refreshToken, apiTld string) *Config {
	var TokenURL = fmt.Sprintf("https://accounts.zoho.%s/oauth/v2/token", apiTld)

	return &Config{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				AuthStyle: oauth2.AuthStyleInParams,
				TokenURL:  TokenURL,
			},
		},
		RefreshToken: refreshToken,
	}
}

// Client returns a *http.Client which automatically retrieves OAuth access
// tokens and attaches them to any request made with it.
func (c *Config) Client(ctx context.Context) *http.Client {
	return &http.Client{
		Transport: &oauth2.Transport{
			Source: c.TokenSource(ctx),
		},
	}
}

// TokenSource creates an oauth2.TokenSource which obtains access tokens using
// the refresh token.
func (c *Config) TokenSource(ctx context.Context) oauth2.TokenSource {
	t := &oauth2.Token{
		RefreshToken: c.RefreshToken,
		TokenType:    TokenType,
	}

	return &tokenSource{
		delegate: c.Config.TokenSource(ctx, t),
	}
}
