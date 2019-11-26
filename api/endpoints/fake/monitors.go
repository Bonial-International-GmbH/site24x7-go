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
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.Monitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Monitors) Create(monitor *api.Monitor) (*api.Monitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.Monitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Monitors) Update(monitor *api.Monitor) (*api.Monitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.Monitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Monitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *Monitors) List() ([]*api.Monitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.Monitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Monitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *Monitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
