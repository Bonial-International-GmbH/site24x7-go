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
	return mockReturn[api.ITAutomation](e.Called(actionID))
}
func (e *ITAutomations) Create(automation *api.ITAutomation) (*api.ITAutomation, error) {
	return mockReturn[api.ITAutomation](e.Called(automation))
}
func (e *ITAutomations) Update(automation *api.ITAutomation) (*api.ITAutomation, error) {
	return mockReturn[api.ITAutomation](e.Called(automation))
}
func (e *ITAutomations) Delete(actionID string) error          { return e.Called(actionID).Error(0) }
func (e *ITAutomations) List() ([]*api.ITAutomation, error)    { return mockReturnSlice[api.ITAutomation](e.Called()) }
