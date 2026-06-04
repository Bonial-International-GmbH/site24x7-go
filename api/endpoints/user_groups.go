package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type UserGroups interface {
	Get(groupID string) (*api.UserGroup, error)
	Create(group *api.UserGroup) (*api.UserGroup, error)
	Update(group *api.UserGroup) (*api.UserGroup, error)
	Delete(groupID string) error
	List() ([]*api.UserGroup, error)
}

type userGroups struct {
	crud[api.UserGroup]
}

func NewUserGroups(client rest.Client) UserGroups {
	return &userGroups{crud[api.UserGroup]{client: client, resource: "user_groups"}}
}

func (c *userGroups) Get(groupID string) (*api.UserGroup, error) { return c.get(groupID) }
func (c *userGroups) Create(group *api.UserGroup) (*api.UserGroup, error) {
	return c.create(group)
}
func (c *userGroups) Update(group *api.UserGroup) (*api.UserGroup, error) {
	return c.update(group.UserGroupID, group)
}
func (c *userGroups) Delete(groupID string) error       { return c.delete(groupID) }
func (c *userGroups) List() ([]*api.UserGroup, error)   { return c.list() }
