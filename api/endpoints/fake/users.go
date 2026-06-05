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

func (e *Users) Get(userID string) (*api.User, error) { return mockReturn[api.User](e.Called(userID)) }
func (e *Users) Create(user *api.User) (*api.User, error) {
	return mockReturn[api.User](e.Called(user))
}
func (e *Users) Update(user *api.User) (*api.User, error) {
	return mockReturn[api.User](e.Called(user))
}
func (e *Users) Delete(userID string) error { return e.Called(userID).Error(0) }
func (e *Users) List() ([]*api.User, error) { return mockReturnSlice[api.User](e.Called()) }
