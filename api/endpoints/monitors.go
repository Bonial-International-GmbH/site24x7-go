package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type Monitors interface {
	Get(monitorID string) (*api.Monitor, error)
	Create(monitor *api.Monitor) (*api.Monitor, error)
	Update(monitor *api.Monitor) (*api.Monitor, error)
	Delete(monitorID string) error
	List() ([]*api.Monitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type monitors struct {
	crud[api.Monitor]
}

func NewMonitors(client rest.Client) Monitors {
	return &monitors{crud[api.Monitor]{client: client, resource: "monitors"}}
}

func (c *monitors) Get(monitorID string) (*api.Monitor, error)        { return c.get(monitorID) }
func (c *monitors) Create(monitor *api.Monitor) (*api.Monitor, error) { return c.create(monitor) }
func (c *monitors) Update(monitor *api.Monitor) (*api.Monitor, error) {
	return c.update(monitor.MonitorID, monitor)
}
func (c *monitors) Delete(monitorID string) error      { return c.delete(monitorID) }
func (c *monitors) List() ([]*api.Monitor, error)      { return c.list() }

func (c *monitors) Activate(monitorID string) error {
	return c.client.Put().Resource("monitors/activate").ResourceID(monitorID).Do().Err()
}

func (c *monitors) Suspend(monitorID string) error {
	return c.client.Put().Resource("monitors/suspend").ResourceID(monitorID).Do().Err()
}
