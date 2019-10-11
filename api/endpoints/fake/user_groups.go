package fake

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.UserGroups = &UserGroups{}

type UserGroups struct {
	mock.Mock
}

func (e *UserGroups) Get(groupID string) (*api.UserGroup, error) {
	args := e.Called(groupID)
	if obj, ok := args.Get(0).(*api.UserGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *UserGroups) Create(group *api.UserGroup) (*api.UserGroup, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.UserGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *UserGroups) Update(group *api.UserGroup) (*api.UserGroup, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.UserGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *UserGroups) Delete(groupID string) error {
	args := e.Called(groupID)
	return args.Error(0)
}

func (e *UserGroups) List() ([]*api.UserGroup, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.UserGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
