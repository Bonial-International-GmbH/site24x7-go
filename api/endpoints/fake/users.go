package fake

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.Users = &Users{}

type Users struct {
	mock.Mock
}

func (e *Users) Get(userID string) (*api.User, error) {
	args := e.Called(userID)
	if obj, ok := args.Get(0).(*api.User); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Users) Create(user *api.User) (*api.User, error) {
	args := e.Called(user)
	if obj, ok := args.Get(0).(*api.User); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Users) Update(user *api.User) (*api.User, error) {
	args := e.Called(user)
	if obj, ok := args.Get(0).(*api.User); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Users) Delete(userID string) error {
	args := e.Called(userID)
	return args.Error(0)
}

func (e *Users) List() ([]*api.User, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.User); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
