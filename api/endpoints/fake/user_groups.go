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
	return mockReturn[api.UserGroup](e.Called(groupID))
}
func (e *UserGroups) Create(group *api.UserGroup) (*api.UserGroup, error) {
	return mockReturn[api.UserGroup](e.Called(group))
}
func (e *UserGroups) Update(group *api.UserGroup) (*api.UserGroup, error) {
	return mockReturn[api.UserGroup](e.Called(group))
}
func (e *UserGroups) Delete(groupID string) error { return e.Called(groupID).Error(0) }
func (e *UserGroups) List() ([]*api.UserGroup, error) {
	return mockReturnSlice[api.UserGroup](e.Called())
}
