package fake

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.Monitors = &Monitors{}

type Monitors struct {
	mock.Mock
}

func (e *Monitors) Get(monitorID string) (*api.Monitor, error) {
	return mockReturn[api.Monitor](e.Called(monitorID))
}
func (e *Monitors) Create(monitor *api.Monitor) (*api.Monitor, error) {
	return mockReturn[api.Monitor](e.Called(monitor))
}
func (e *Monitors) Update(monitor *api.Monitor) (*api.Monitor, error) {
	return mockReturn[api.Monitor](e.Called(monitor))
}
func (e *Monitors) List() ([]*api.Monitor, error)   { return mockReturnSlice[api.Monitor](e.Called()) }
func (e *Monitors) Delete(monitorID string) error   { return e.Called(monitorID).Error(0) }
func (e *Monitors) Activate(monitorID string) error { return e.Called(monitorID).Error(0) }
func (e *Monitors) Suspend(monitorID string) error  { return e.Called(monitorID).Error(0) }
