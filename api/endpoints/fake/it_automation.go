package fake

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.ITAutomations = &ITAutomations{}

type ITAutomations struct {
	mock.Mock
}

func (e *ITAutomations) Get(actionID string) (*api.ITAutomation, error) {
	args := e.Called(actionID)
	if obj, ok := args.Get(0).(*api.ITAutomation); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ITAutomations) Create(automation *api.ITAutomation) (*api.ITAutomation, error) {
	args := e.Called(automation)
	if obj, ok := args.Get(0).(*api.ITAutomation); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ITAutomations) Update(automation *api.ITAutomation) (*api.ITAutomation, error) {
	args := e.Called(automation)
	if obj, ok := args.Get(0).(*api.ITAutomation); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ITAutomations) Delete(actionID string) error {
	args := e.Called(actionID)
	return args.Error(0)
}

func (e *ITAutomations) List() ([]*api.ITAutomation, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.ITAutomation); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
