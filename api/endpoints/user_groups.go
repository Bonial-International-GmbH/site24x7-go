package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type UserGroupsEndpoint interface {
	Get(groupID string) (*api.UserGroup, error)
	Create(group *api.UserGroup) (*api.UserGroup, error)
	Update(group *api.UserGroup) (*api.UserGroup, error)
	Delete(groupID string) error
	List() ([]*api.UserGroup, error)
}

type userGroupsEndpoint struct {
	client rest.Client
}

// TODO(mohmann) implement UserGroupsEndpoint interface here
