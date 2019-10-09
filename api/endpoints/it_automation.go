package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type ITAutomationsEndpoint interface {
	Get(actionID string) (*api.ITAutomations, error)
	Create(automation *api.ITAutomations) (*api.ITAutomations, error)
	Update(automation *api.ITAutomations) (*api.ITAutomations, error)
	Delete(actionID string) error
	List() ([]*api.ITAutomations, error)
}

type itAutomationsEndpoint struct {
	client rest.Client
}

func NewITAutomationsEndpoint(client rest.Client) ITAutomationsEndpoint {
	return &itAutomationsEndpoint{
		client: client,
	}
}

func (c *itAutomationsEndpoint) Get(actionID string) (*api.ITAutomations, error) {
	automation := &api.ITAutomations{}
	err := c.client.
		Get().
		Resource("it_automation").
		ResourceID(actionID).
		Do().
		Into(automation)

	return automation, err
}

func (c *itAutomationsEndpoint) Create(automation *api.ITAutomations) (*api.ITAutomations, error) {
	newITAutomation := &api.ITAutomations{}
	err := c.client.
		Post().
		Resource("it_automation").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(automation).
		Do().
		Into(newITAutomation)

	return newITAutomation, err
}

func (c *itAutomationsEndpoint) Update(automation *api.ITAutomations) (*api.ITAutomations, error) {
	itAutomation := &api.ITAutomations{}
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

func (c *itAutomationsEndpoint) List() ([]*api.ITAutomations, error) {
	itAutomation := []*api.ITAutomations{}
	err := c.client.
		Get().
		Resource("it_automation").
		Do().
		Into(&itAutomation)

	return itAutomation, err
}
