package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type GroupsEndpoint interface {
	Get(groupID string) (*api.Group, error)
	Create(group *api.Group) (*api.Group, error)
	Update(group *api.Group) (*api.Group, error)
	Delete(groupID string) error
	List() ([]*api.Group, error)
}

type groupsEndpoint struct {
	client rest.Client
}

func NewGroupsEndpoint(client rest.Client) GroupsEndpoint {
	return &groupsEndpoint{
		client: client,
	}
}

func (c *groupsEndpoint) Get(groupID string) (*api.Group, error) {
	group := &api.Group{}
	err := c.client.
		Get().
		Resource("monitor_groups").
		ResourceID(groupID).
		Do().
		Into(group)

	return group, err
}

func (c *groupsEndpoint) Create(group *api.Group) (*api.Group, error) {
	newGroup := &api.Group{}
	err := c.client.
		Post().
		Resource("monitor_groups").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Into(newGroup)

	return newGroup, err
}

func (c *groupsEndpoint) Update(group *api.Group) (*api.Group, error) {
	updatedGroup := &api.Group{}
	err := c.client.
		Put().
		Resource("monitor_groups").
		ResourceID(group.GroupID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Into(updatedGroup)

	return updatedGroup, err
}

func (c *groupsEndpoint) Delete(groupID string) error {
	return c.client.
		Delete().
		Resource("monitor_groups").
		ResourceID(groupID).
		Do().
		Err()
}

func (c *groupsEndpoint) List() ([]*api.Group, error) {
	groups := []*api.Group{}
	err := c.client.
		Get().
		Resource("monitor_groups").
		Do().
		Into(&groups)

	return groups, err
}
