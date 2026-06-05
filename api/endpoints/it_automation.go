package endpoints

import (
	"github.com/Bonial-International-GmbH/site24x7-go/api"
	"github.com/Bonial-International-GmbH/site24x7-go/rest"
)

type ITAutomations interface {
	Get(actionID string) (*api.ITAutomation, error)
	Create(automation *api.ITAutomation) (*api.ITAutomation, error)
	Update(automation *api.ITAutomation) (*api.ITAutomation, error)
	Delete(actionID string) error
	List() ([]*api.ITAutomation, error)
}

type itAutomations struct {
	crud[api.ITAutomation]
}

func NewITAutomations(client rest.Client) ITAutomations {
	return &itAutomations{crud[api.ITAutomation]{client: client, resource: "it_automation"}}
}

func (c *itAutomations) Get(actionID string) (*api.ITAutomation, error) { return c.get(actionID) }
func (c *itAutomations) Create(automation *api.ITAutomation) (*api.ITAutomation, error) {
	return c.create(automation)
}
func (c *itAutomations) Update(automation *api.ITAutomation) (*api.ITAutomation, error) {
	return c.update(automation.ActionID, automation)
}
func (c *itAutomations) Delete(actionID string) error       { return c.delete(actionID) }
func (c *itAutomations) List() ([]*api.ITAutomation, error) { return c.list() }
