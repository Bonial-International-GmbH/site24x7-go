package fake

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.LocationProfiles = &LocationProfiles{}

type LocationProfiles struct {
	mock.Mock
}

func (e *LocationProfiles) Get(profileID string) (*api.LocationProfile, error) {
	return mockReturn[api.LocationProfile](e.Called(profileID))
}
func (e *LocationProfiles) Create(profile *api.LocationProfile) (*api.LocationProfile, error) {
	return mockReturn[api.LocationProfile](e.Called(profile))
}
func (e *LocationProfiles) Update(profile *api.LocationProfile) (*api.LocationProfile, error) {
	return mockReturn[api.LocationProfile](e.Called(profile))
}
func (e *LocationProfiles) Delete(profileID string) error { return e.Called(profileID).Error(0) }
func (e *LocationProfiles) List() ([]*api.LocationProfile, error) {
	return mockReturnSlice[api.LocationProfile](e.Called())
}
