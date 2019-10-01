package site24x7

import (
	"context"
	"net/http"

	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/Bonial-International-GmbH/site24x7-go/backoff"
	"github.com/Bonial-International-GmbH/site24x7-go/oauth"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

// Config is the configuration for the Site24x7 API Client.
type Config struct {
	// ClientID is the OAuth client ID needed to obtain an access token for API
	// usage.
	ClientID string

	// ClientSecret is the OAuth client secret needed to obtain an access token
	// for API usage.
	ClientSecret string

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string

	// RetryConfig contains the configuration of the backoff-retry behavior. If
	// nil, backoff.DefaultRetryConfig will be used.
	RetryConfig *backoff.RetryConfig
}

// OAuthClient creates a new *http.Client from c that transparently obtains and
// attaches OAuth access tokens to every request.
func (c *Config) OAuthClient(ctx context.Context) *http.Client {
	oauthConfig := oauth.NewConfig(c.ClientID, c.ClientSecret, c.RefreshToken)

	return oauthConfig.Client(ctx)
}

// HTTPClient is the interface of an http client that is compatible with
// *http.Client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the Site24x7 API Client interface. It provides methods to get
// clients for resource endpoints.
type Client interface {
	LocationProfiles() endpoints.LocationProfilesEndpoint
	MonitorGroups() endpoints.MonitorGroupsEndpoint
	Monitors() endpoints.MonitorsEndpoint
	NotificationProfiles() endpoints.NotificationProfilesEndpoint
	ThresholdProfiles() endpoints.ThresholdProfilesEndpoint
	UserGroups() endpoints.UserGroupsEndpoint
}

type client struct {
	restClient rest.Client
}

// New creates a new Site24x7 API Client with Config c.
func New(c Config) Client {
	httpClient := backoff.WithRetries(
		c.OAuthClient(context.Background()),
		c.RetryConfig,
	)

	return NewClient(httpClient)
}

// NewClient creates a new Site24x7 API Client from httpClient. This can be
// used to provide a custom http client for use with the API. The custom http
// client has to transparently handle the Site24x7 OAuth flow.
func NewClient(httpClient HTTPClient) Client {
	return &client{
		restClient: rest.NewClient(httpClient),
	}
}

// LocationProfiles implements Client.
func (c *client) LocationProfiles() endpoints.LocationProfilesEndpoint {
	return endpoints.NewLocationProfilesEndpoint(c.restClient)
}

// Monitors implements Client.
func (c *client) Monitors() endpoints.MonitorsEndpoint {
	return endpoints.NewMonitorsEndpoint(c.restClient)
}

// MonitorGroups implements Client.
func (c *client) MonitorGroups() endpoints.MonitorGroupsEndpoint {
	return endpoints.NewMonitorGroupsEndpoint(c.restClient)
}

// NotificationProfiles implements Client.
func (c *client) NotificationProfiles() endpoints.NotificationProfilesEndpoint {
	return endpoints.NewNotificationProfilesEndpoint(c.restClient)
}

// ThresholdProfiles implements Client.
func (c *client) ThresholdProfiles() endpoints.ThresholdProfilesEndpoint {
	return endpoints.NewThresholdProfilesEndpoint(c.restClient)
}

// UserGroups implements Client.
func (c *client) UserGroups() endpoints.UserGroupsEndpoint {
	// TODO(mohmann) use userGroupsEndpoint here
	return nil
}
