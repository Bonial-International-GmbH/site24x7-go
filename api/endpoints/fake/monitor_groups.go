package fake

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.MonitorGroups = &MonitorGroups{}

type MonitorGroups struct {
	mock.Mock
}

func (e *MonitorGroups) Get(groupID string) (*api.MonitorGroup, error) {
	return mockReturn[api.MonitorGroup](e.Called(groupID))
}
func (e *MonitorGroups) Create(group *api.MonitorGroup) (*api.MonitorGroup, error) {
	return mockReturn[api.MonitorGroup](e.Called(group))
}
func (e *MonitorGroups) Update(group *api.MonitorGroup) (*api.MonitorGroup, error) {
	return mockReturn[api.MonitorGroup](e.Called(group))
}
func (e *MonitorGroups) Delete(groupID string) error { return e.Called(groupID).Error(0) }
func (e *MonitorGroups) List() ([]*api.MonitorGroup, error) {
	return mockReturnSlice[api.MonitorGroup](e.Called())
}
