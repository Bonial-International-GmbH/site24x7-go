package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type NotificationProfilesEndpoint interface {
	Get(profileID string) (*api.NotificationProfile, error)
	Create(profile *api.NotificationProfile) (*api.NotificationProfile, error)
	Update(profile *api.NotificationProfile) (*api.NotificationProfile, error)
	Delete(profileID string) error
	List() ([]*api.NotificationProfile, error)
}

type notificationProfilesEndpoint struct {
	client rest.Client
}

// TODO(mohmann) implement NotificationProfilesEndpoint interface here
