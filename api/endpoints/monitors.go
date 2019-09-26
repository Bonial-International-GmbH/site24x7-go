package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type MonitorsEndpoint interface {
	Get(monitorID string) (*api.Monitor, error)
	Create(monitor *api.Monitor) (*api.Monitor, error)
	Update(monitor *api.Monitor) (*api.Monitor, error)
	Delete(monitorID string) error
	List() ([]*api.Monitor, error)
}

type monitorsEndpoint struct {
	client rest.Client
}

func NewMonitorsEndpoint(client rest.Client) MonitorsEndpoint {
	return &monitorsEndpoint{
		client: client,
	}
}

func (c *monitorsEndpoint) Get(monitorID string) (*api.Monitor, error) {
	monitor := &api.Monitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Into(monitor)

	return monitor, err
}

func (c *monitorsEndpoint) Create(monitor *api.Monitor) (*api.Monitor, error) {
	newMonitor := &api.Monitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Into(newMonitor)

	return newMonitor, err
}

func (c *monitorsEndpoint) Update(monitor *api.Monitor) (*api.Monitor, error) {
	updatedMonitor := &api.Monitor{}
	err := c.client.
		Put().
		Resource("monitors").
		ResourceID(monitor.MonitorID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Into(updatedMonitor)

	return updatedMonitor, err
}

func (c *monitorsEndpoint) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *monitorsEndpoint) List() ([]*api.Monitor, error) {
	monitors := []*api.Monitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Into(&monitors)

	return monitors, err
}
