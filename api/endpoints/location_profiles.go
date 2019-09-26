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

// TODO(mohmann) implement LocationProfilesEndpoint interface here
