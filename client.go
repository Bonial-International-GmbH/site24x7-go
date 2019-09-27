package site24x7

import (
	"context"
	"net/http"

	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
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

// NewClientForConfig creates a new Site24x7 API Client for Config c.
func NewClientForConfig(c Config) Client {
	oauthConfig := oauth.NewConfig(c.ClientID, c.ClientSecret, c.RefreshToken)

	httpClient := oauthConfig.Client(context.Background())

	return NewClient(httpClient)
}

// NewClient creates a new Site24x7 API Client from httpClient.
func NewClient(httpClient *http.Client) Client {
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
	// TODO(mohmann) use thresholdProfilesEndpoint here
	return nil
}

// UserGroups implements Client.
func (c *client) UserGroups() endpoints.UserGroupsEndpoint {
	// TODO(mohmann) use userGroupsEndpoint here
	return nil
}
