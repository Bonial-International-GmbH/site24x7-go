package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type ThresholdProfiles interface {
	Get(profileID string) (*api.ThresholdProfile, error)
	Create(profile *api.ThresholdProfile) (*api.ThresholdProfile, error)
	Update(profile *api.ThresholdProfile) (*api.ThresholdProfile, error)
	Delete(profileID string) error
	List() ([]*api.ThresholdProfile, error)
}

type thresholdProfiles struct {
	crud[api.ThresholdProfile]
}

func NewThresholdProfiles(client rest.Client) ThresholdProfiles {
	return &thresholdProfiles{crud[api.ThresholdProfile]{client: client, resource: "threshold_profiles"}}
}

func (c *thresholdProfiles) Get(profileID string) (*api.ThresholdProfile, error) {
	return c.get(profileID)
}
func (c *thresholdProfiles) Create(profile *api.ThresholdProfile) (*api.ThresholdProfile, error) {
	return c.create(profile)
}
func (c *thresholdProfiles) Update(profile *api.ThresholdProfile) (*api.ThresholdProfile, error) {
	return c.update(profile.ProfileID, profile)
}
func (c *thresholdProfiles) Delete(profileID string) error          { return c.delete(profileID) }
func (c *thresholdProfiles) List() ([]*api.ThresholdProfile, error) { return c.list() }
