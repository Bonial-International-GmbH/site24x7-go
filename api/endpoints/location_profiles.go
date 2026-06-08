package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type LocationProfiles interface {
	Get(profileID string) (*api.LocationProfile, error)
	Create(profile *api.LocationProfile) (*api.LocationProfile, error)
	Update(profile *api.LocationProfile) (*api.LocationProfile, error)
	Delete(profileID string) error
	List() ([]*api.LocationProfile, error)
}

type locationProfiles struct {
	crud[api.LocationProfile]
}

func NewLocationProfiles(client rest.Client) LocationProfiles {
	return &locationProfiles{crud[api.LocationProfile]{client: client, resource: "location_profiles"}}
}

func (c *locationProfiles) Get(profileID string) (*api.LocationProfile, error) {
	return c.get(profileID)
}
func (c *locationProfiles) Create(profile *api.LocationProfile) (*api.LocationProfile, error) {
	return c.create(profile)
}
func (c *locationProfiles) Update(profile *api.LocationProfile) (*api.LocationProfile, error) {
	return c.update(profile.ProfileID, profile)
}
func (c *locationProfiles) Delete(profileID string) error         { return c.delete(profileID) }
func (c *locationProfiles) List() ([]*api.LocationProfile, error) { return c.list() }
