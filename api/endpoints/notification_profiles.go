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

// func NewNotificationProfilesEndpoint(client rest.Client) NotificationProfilesEndpoint {
// 	return &notificationProfilesEndpoint{
// 		client: client,
// 	}
// }

// func (c *notificationProfilesEndpoint) Get(profileID string) (*api.NotificationProfile, error) {
// 	profile := &api.NotificationProfile{}
// 	err := c.client.
// 		Get().
// 		Resource("notification_profiles").
// 		ResourceID(profileID).
// 		Do().
// 		Into(profile)

// 	return profile, err
// }
