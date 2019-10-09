package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type ITAutomationsEndpoint interface {
	Get(actionID string) (*api.ITAutomation, error)
	Create(automation *api.ITAutomation) (*api.ITAutomation, error)
	Update(automation *api.ITAutomation) (*api.ITAutomation, error)
	Delete(actionID string) error
	List() ([]*api.ITAutomation, error)
}

type itAutomationsEndpoint struct {
	client rest.Client
}

func NewITAutomationsEndpoint(client rest.Client) ITAutomationsEndpoint {
	return &itAutomationsEndpoint{
		client: client,
	}
}

func (c *itAutomationsEndpoint) Get(actionID string) (*api.ITAutomation, error) {
	automation := &api.ITAutomation{}
	err := c.client.
		Get().
		Resource("it_automation").
		ResourceID(actionID).
		Do().
		Into(automation)

	return automation, err
}

func (c *itAutomationsEndpoint) Create(automation *api.ITAutomation) (*api.ITAutomation, error) {
	newITAutomation := &api.ITAutomation{}
	err := c.client.
		Post().
		Resource("it_automation").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(automation).
		Do().
		Into(newITAutomation)

	return newITAutomation, err
}

func (c *itAutomationsEndpoint) Update(automation *api.ITAutomation) (*api.ITAutomation, error) {
	itAutomation := &api.ITAutomation{}
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

func (c *itAutomationsEndpoint) Delete(actionID string) error {
	return c.client.
		Delete().
		Resource("it_automation").
		ResourceID(actionID).
		Do().
		Err()
}

func (c *itAutomationsEndpoint) List() ([]*api.ITAutomation, error) {
	itAutomation := []*api.ITAutomation{}
	err := c.client.
		Get().
		Resource("it_automation").
		Do().
		Into(&itAutomation)

	return itAutomation, err
}
