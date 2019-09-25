package api

import (
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RequestFactory interface {
	Request() *Request
}

type Client interface {
	LocationProfiles() LocationProfilesEndpoint
	MonitorGroups() MonitorGroupsEndpoint
	Monitors() MonitorsEndpoint
	NotificationProfiles() NotificationProfilesEndpoint
	ThresholdProfiles() ThresholdProfilesEndpoint
	UserGroups() UserGroupsEndpoint
}

type client struct {
	httpClient HTTPClient
}

func NewClient(httpClient HTTPClient) Client {
	return &client{
		httpClient: httpClient,
	}
}

func (c *client) LocationProfiles() LocationProfilesEndpoint {
	// TODO(mohmann) use locationProfilesEndpoint here
	return nil
}

func (c *client) Monitors() MonitorsEndpoint {
	return &monitorsEndpoint{client: c}
}

func (c *client) MonitorGroups() MonitorGroupsEndpoint {
	// TODO(mohmann) use monitorGroupsEndpoint here
	return nil
}

func (c *client) NotificationProfiles() NotificationProfilesEndpoint {
	// TODO(mohmann) use notificationProfilesEndpoint here
	return nil
}

func (c *client) ThresholdProfiles() ThresholdProfilesEndpoint {
	// TODO(mohmann) use thresholdProfilesEndpoint here
	return nil
}

func (c *client) UserGroups() UserGroupsEndpoint {
	// TODO(mohmann) use userGroupsEndpoint here
	return nil
}

func (c *client) Request() *Request {
	req := NewRequest(c.httpClient).
		AddHeader("Accept", "application/json; version=2.0")

	return req
}
