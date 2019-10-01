package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type LocationProfilesEndpoint interface {
	Get(profileID string) (*api.LocationProfile, error)
	Create(profile *api.LocationProfile) (*api.LocationProfile, error)
	Update(profile *api.LocationProfile) (*api.LocationProfile, error)
	Delete(profileID string) error
	List() ([]*api.LocationProfile, error)
}

type locationProfilesEndpoint struct {
	client rest.Client
}

func NewLocationProfilesEndpoint(client rest.Client) LocationProfilesEndpoint {
	return &locationProfilesEndpoint{
		client: client,
	}
}

func (c *locationProfilesEndpoint) Get(profileID string) (*api.LocationProfile, error) {
	profile := &api.LocationProfile{}
	err := c.client.
		Get().
		Resource("location_profiles").
		ResourceID(profileID).
		Do().
		Into(profile)

	return profile, err
}

func (c *locationProfilesEndpoint) Create(profile *api.LocationProfile) (*api.LocationProfile, error) {
	newProfile := &api.LocationProfile{}
	err := c.client.
		Post().
		Resource("location_profiles").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Into(newProfile)

	return newProfile, err
}

func (c *locationProfilesEndpoint) Update(profile *api.LocationProfile) (*api.LocationProfile, error) {
	updatedProfile := &api.LocationProfile{}
	err := c.client.
		Put().
		Resource("location_profiles").
		ResourceID(profile.ProfileID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Into(updatedProfile)

	return updatedProfile, err
}

func (c *locationProfilesEndpoint) Delete(profileID string) error {
	return c.client.
		Delete().
		Resource("location_profiles").
		ResourceID(profileID).
		Do().
		Err()
}

func (c *locationProfilesEndpoint) List() ([]*api.LocationProfile, error) {
	profiles := []*api.LocationProfile{}
	err := c.client.
		Get().
		Resource("location_profiles").
		Do().
		Into(&profiles)

	return profiles, err
}
