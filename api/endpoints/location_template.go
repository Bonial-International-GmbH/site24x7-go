package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type LocationTemplate interface {
	Get() (*api.LocationTemplate, error)
}

type locationTemplate struct {
	client rest.Client
}

func NewLocationTemplate(client rest.Client) LocationTemplate {
	return &locationTemplate{
		client: client,
	}
}

func (c *locationTemplate) Get() (*api.LocationTemplate, error) {
	template := &api.LocationTemplate{}
	err := c.client.
		Get().
		Resource("location_template").
		Do().
		Into(&template)

	return template, err
}
