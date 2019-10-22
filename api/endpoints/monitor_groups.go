package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type MonitorGroups interface {
	Get(groupID string) (*api.MonitorGroup, error)
	Create(group *api.MonitorGroup) (*api.MonitorGroup, error)
	Update(group *api.MonitorGroup) (*api.MonitorGroup, error)
	Delete(groupID string) error
	List() ([]*api.MonitorGroup, error)
}

type monitorGroups struct {
	client rest.Client
}

func NewMonitorGroups(client rest.Client) MonitorGroups {
	return &monitorGroups{
		client: client,
	}
}

func (c *monitorGroups) Get(groupID string) (*api.MonitorGroup, error) {
	monitorGroup := &api.MonitorGroup{}
	err := c.client.
		Get().
		Resource("monitor_groups").
		ResourceID(groupID).
		Do().
		Into(monitorGroup)

	return monitorGroup, err
}

func (c *monitorGroups) Create(group *api.MonitorGroup) (*api.MonitorGroup, error) {
	newMonitorGroup := &api.MonitorGroup{}
	err := c.client.
		Post().
		Resource("monitor_groups").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Into(newMonitorGroup)

	return newMonitorGroup, err
}

func (c *monitorGroups) Update(group *api.MonitorGroup) (*api.MonitorGroup, error) {
	updatedGroup := &api.MonitorGroup{}
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

func (c *monitorGroups) Delete(groupID string) error {
	return c.client.
		Delete().
		Resource("monitor_groups").
		ResourceID(groupID).
		Do().
		Err()
}

func (c *monitorGroups) List() ([]*api.MonitorGroup, error) {
	monitorGroups := []*api.MonitorGroup{}
	err := c.client.
		Get().
		Resource("monitor_groups").
		Do().
		Into(&monitorGroups)

	return monitorGroups, err
}
