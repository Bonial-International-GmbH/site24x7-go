package api

type MonitorsEndpoint interface {
	Get(monitorID string) (*Monitor, error)
	Create(monitor *Monitor) (*Monitor, error)
	Update(monitor *Monitor) (*Monitor, error)
	Delete(monitorID string) error
	List() ([]*Monitor, error)
}

type monitorsEndpoint struct {
	client RequestFactory
}

func (c *monitorsEndpoint) Get(monitorID string) (*Monitor, error) {
	monitor := &Monitor{}
	err := c.client.
		Request().
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Into(monitor)

	return monitor, err
}

func (c *monitorsEndpoint) Create(monitor *Monitor) (*Monitor, error) {
	newMonitor := &Monitor{}
	err := c.client.
		Request().
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Into(newMonitor)

	return newMonitor, err
}

func (c *monitorsEndpoint) Update(monitor *Monitor) (*Monitor, error) {
	updatedMonitor := &Monitor{}
	err := c.client.
		Request().
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
		Request().
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *monitorsEndpoint) List() ([]*Monitor, error) {
	monitors := []*Monitor{}
	err := c.client.
		Request().
		Get().
		Resource("monitors").
		Do().
		Into(&monitors)

	return monitors, err
}
