package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type NotificationProfiles interface {
	Get(profileID string) (*api.NotificationProfile, error)
	Create(profile *api.NotificationProfile) (*api.NotificationProfile, error)
	Update(profile *api.NotificationProfile) (*api.NotificationProfile, error)
	Delete(profileID string) error
	List() ([]*api.NotificationProfile, error)
}

type notificationProfiles struct {
	crud[api.NotificationProfile]
}

func NewNotificationProfiles(client rest.Client) NotificationProfiles {
	return &notificationProfiles{crud[api.NotificationProfile]{client: client, resource: "notification_profiles"}}
}

func (c *notificationProfiles) Get(profileID string) (*api.NotificationProfile, error) {
	return c.get(profileID)
}
func (c *notificationProfiles) Create(profile *api.NotificationProfile) (*api.NotificationProfile, error) {
	return c.create(profile)
}
func (c *notificationProfiles) Update(profile *api.NotificationProfile) (*api.NotificationProfile, error) {
	return c.update(profile.ProfileID, profile)
}
func (c *notificationProfiles) Delete(profileID string) error              { return c.delete(profileID) }
func (c *notificationProfiles) List() ([]*api.NotificationProfile, error)  { return c.list() }
