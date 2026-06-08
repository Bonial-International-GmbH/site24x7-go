package fake

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.ThresholdProfiles = &ThresholdProfiles{}

type ThresholdProfiles struct {
	mock.Mock
}

func (e *ThresholdProfiles) Get(profileID string) (*api.ThresholdProfile, error) {
	return mockReturn[api.ThresholdProfile](e.Called(profileID))
}
func (e *ThresholdProfiles) Create(profile *api.ThresholdProfile) (*api.ThresholdProfile, error) {
	return mockReturn[api.ThresholdProfile](e.Called(profile))
}
func (e *ThresholdProfiles) Update(profile *api.ThresholdProfile) (*api.ThresholdProfile, error) {
	return mockReturn[api.ThresholdProfile](e.Called(profile))
}
func (e *ThresholdProfiles) Delete(profileID string) error { return e.Called(profileID).Error(0) }
func (e *ThresholdProfiles) List() ([]*api.ThresholdProfile, error) {
	return mockReturnSlice[api.ThresholdProfile](e.Called())
}
