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

func NewNotificationProfilesEndpoint(client rest.Client) NotificationProfilesEndpoint {
	return &notificationProfilesEndpoint{
		client: client,
	}
}

func (c *notificationProfilesEndpoint) Get(profileID string) (*api.NotificationProfile, error) {
	profile := &api.NotificationProfile{}
	err := c.client.
		Get().
		Resource("notification_profiles").
		ResourceID(profileID).
		Do().
		Into(profile)

	return profile, err
}

func (c *notificationProfilesEndpoint) Create(profile *api.NotificationProfile) (*api.NotificationProfile, error) {
	newNotificationProfile := &api.NotificationProfile{}
	err := c.client.
		Post().
		Resource("notification_profiles").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Into(newNotificationProfile)

	return newNotificationProfile, err
}

func (c *notificationProfilesEndpoint) Update(profile *api.NotificationProfile) (*api.NotificationProfile, error) {
	updatedNotificationProfile := &api.NotificationProfile{}
	err := c.client.
		Put().
		Resource("notification_profiles").
		ResourceID(profile.ProfileID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Into(updatedNotificationProfile)

	return updatedNotificationProfile, err
}

func (c *notificationProfilesEndpoint) Delete(profileID string) error {
	return c.client.
		Delete().
		Resource("notification_profiles").
		ResourceID(profileID).
		Do().
		Err()
}

func (c *notificationProfilesEndpoint) List() ([]*api.NotificationProfile, error) {
	notificationProfiles := []*api.NotificationProfile{}
	err := c.client.
		Get().
		Resource("notification_profiles").
		Do().
		Into(&notificationProfiles)

	return notificationProfiles, err
}
