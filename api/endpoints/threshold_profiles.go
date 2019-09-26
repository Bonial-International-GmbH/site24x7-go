package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type ThresholdProfilesEndpoint interface {
	Get(profileID string) (*api.ThresholdProfile, error)
	Create(profile *api.ThresholdProfile) (*api.ThresholdProfile, error)
	Update(profile *api.ThresholdProfile) (*api.ThresholdProfile, error)
	Delete(profileID string) error
	List() ([]*api.ThresholdProfile, error)
}

type thresholdProfilesEndpoint struct {
	client rest.Client
}

// TODO(mohmann) implement ThresholdProfilesEndpoint interface here
