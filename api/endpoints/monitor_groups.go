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
	crud[api.MonitorGroup]
}

func NewMonitorGroups(client rest.Client) MonitorGroups {
	return &monitorGroups{crud[api.MonitorGroup]{client: client, resource: "monitor_groups"}}
}

func (c *monitorGroups) Get(groupID string) (*api.MonitorGroup, error) { return c.get(groupID) }
func (c *monitorGroups) Create(group *api.MonitorGroup) (*api.MonitorGroup, error) {
	return c.create(group)
}
func (c *monitorGroups) Update(group *api.MonitorGroup) (*api.MonitorGroup, error) {
	return c.update(group.GroupID, group)
}
func (c *monitorGroups) Delete(groupID string) error        { return c.delete(groupID) }
func (c *monitorGroups) List() ([]*api.MonitorGroup, error) { return c.list() }
