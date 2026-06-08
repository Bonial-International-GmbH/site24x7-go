package fake

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.NotificationProfiles = &NotificationProfiles{}

type NotificationProfiles struct {
	mock.Mock
}

func (e *NotificationProfiles) Get(profileID string) (*api.NotificationProfile, error) {
	return mockReturn[api.NotificationProfile](e.Called(profileID))
}
func (e *NotificationProfiles) Create(profile *api.NotificationProfile) (*api.NotificationProfile, error) {
	return mockReturn[api.NotificationProfile](e.Called(profile))
}
func (e *NotificationProfiles) Update(profile *api.NotificationProfile) (*api.NotificationProfile, error) {
	return mockReturn[api.NotificationProfile](e.Called(profile))
}
func (e *NotificationProfiles) Delete(profileID string) error { return e.Called(profileID).Error(0) }
func (e *NotificationProfiles) List() ([]*api.NotificationProfile, error) {
	return mockReturnSlice[api.NotificationProfile](e.Called())
}
