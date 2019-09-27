package site24x7

import (
	"net/http"

	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

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

func NewClient(httpClient *http.Client) Client {
	return &client{
		restClient: rest.NewClient(httpClient),
	}
}

func (c *client) LocationProfiles() endpoints.LocationProfilesEndpoint {
	return endpoints.NewLocationProfilesEndpoint(c.restClient)
}

func (c *client) Monitors() endpoints.MonitorsEndpoint {
	return endpoints.NewMonitorsEndpoint(c.restClient)
}

func (c *client) MonitorGroups() endpoints.MonitorGroupsEndpoint {
	return endpoints.NewMonitorGroupsEndpoint(c.restClient)
}

func (c *client) NotificationProfiles() endpoints.NotificationProfilesEndpoint {
	// TODO(mohmann) use notificationProfilesEndpoint here
	return nil
}

func (c *client) ThresholdProfiles() endpoints.ThresholdProfilesEndpoint {
	// TODO(mohmann) use thresholdProfilesEndpoint here
	return nil
}

func (c *client) UserGroups() endpoints.UserGroupsEndpoint {
	// TODO(mohmann) use userGroupsEndpoint here
	return nil
}
