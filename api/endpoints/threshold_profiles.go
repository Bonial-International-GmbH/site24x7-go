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
	client rest.Client
}

func NewThresholdProfiles(client rest.Client) ThresholdProfiles {
	return &thresholdProfiles{
		client: client,
	}
}

func (c *thresholdProfiles) Get(profileID string) (*api.ThresholdProfile, error) {
	profile := &api.ThresholdProfile{}
	err := c.client.
		Get().
		Resource("threshold_profiles").
		ResourceID(profileID).
		Do().
		Into(profile)

	return profile, err
}

func (c *thresholdProfiles) Create(profile *api.ThresholdProfile) (*api.ThresholdProfile, error) {
	newThresholdProfile := &api.ThresholdProfile{}
	err := c.client.
		Post().
		Resource("threshold_profiles").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Into(newThresholdProfile)

	return newThresholdProfile, err
}

func (c *thresholdProfiles) Update(profile *api.ThresholdProfile) (*api.ThresholdProfile, error) {
	updatedThresholdProfile := &api.ThresholdProfile{}
	err := c.client.
		Put().
		Resource("threshold_profiles").
		ResourceID(profile.ProfileID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Into(updatedThresholdProfile)

	return updatedThresholdProfile, err
}

func (c *thresholdProfiles) Delete(profileID string) error {
	return c.client.
		Delete().
		Resource("threshold_profiles").
		ResourceID(profileID).
		Do().
		Err()
}

func (c *thresholdProfiles) List() ([]*api.ThresholdProfile, error) {
	thresholdProfiles := []*api.ThresholdProfile{}
	err := c.client.
		Get().
		Resource("threshold_profiles").
		Do().
		Into(&thresholdProfiles)

	return thresholdProfiles, err
}
