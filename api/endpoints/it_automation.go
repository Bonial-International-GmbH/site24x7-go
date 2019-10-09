package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type ItAutomationEndpoint interface {
	Get(actionID string) (*api.ItAutomation, error)
	Create(automation *api.ItAutomation) (*api.ItAutomation, error)
	Update(automation *api.ItAutomation) (*api.ItAutomation, error)
	Delete(actionID string) error
	List() ([]*api.ItAutomation, error)
}

type itAutomationEndpoint struct {
	client rest.Client
}

func NewItAutomationEndpoint(client rest.Client) ItAutomationEndpoint {
	return &itAutomationEndpoint{
		client: client,
	}
}

func (c *itAutomationEndpoint) Get(actionID string) (*api.ItAutomation, error) {
	automation := &api.ItAutomation{}
	err := c.client.
		Get().
		Resource("it_automation").
		ResourceID(actionID).
		Do().
		Into(automation)

	return automation, err
}

func (c *itAutomationEndpoint) Create(automation *api.ItAutomation) (*api.ItAutomation, error) {
	newItAutomation := &api.ItAutomation{}
	err := c.client.
		Post().
		Resource("it_automation").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(automation).
		Do().
		Into(newItAutomation)

	return newItAutomation, err
}

func (c *itAutomationEndpoint) Update(automation *api.ItAutomation) (*api.ItAutomation, error) {
	itAutomation := &api.ItAutomation{}
	err := c.client.
		Put().
		Resource("it_automation").
		ResourceID(automation.ActionID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(automation).
		Do().
		Into(itAutomation)

	return itAutomation, err
}

func (c *itAutomationEndpoint) Delete(actionID string) error {
	return c.client.
		Delete().
		Resource("it_automation").
		ResourceID(actionID).
		Do().
		Err()
}

func (c *itAutomationEndpoint) List() ([]*api.ItAutomation, error) {
	itAutomation := []*api.ItAutomation{}
	err := c.client.
		Get().
		Resource("it_automation").
		Do().
		Into(&itAutomation)

	return itAutomation, err
}
